package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BetterSearch(slice []int, value int) (bool, int) {
	low := 0
	high := len(slice)

	for low < high {
		mid := (low + high) / 2
		if slice[mid] == value {
			return true, mid
		}
		if slice[mid] < value {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return false, low
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	line := scanner.Text()
	line = strings.Trim(line, "[] ")
	parts := strings.Fields(line)

	slice := []int{}
	for _, p := range parts {
		val, _ := strconv.Atoi(p)
		slice = append(slice, val)
	}

	if !scanner.Scan() {
		return
	}
	searchValue, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	found, result := BetterSearch(slice, searchValue)

	if found {
		fmt.Printf("V %d\n", result)
	} else {
		fmt.Printf("F %d\n", result)
	}
}