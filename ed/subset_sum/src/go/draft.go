package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func backtrack(arr []int, target int, index int) bool {
	if target == 0 {
		return true
	}
	if target < 0 || index >= len(arr) {
		return false
	}
	return backtrack(arr, target-arr[index], index+1) ||
		backtrack(arr, target, index+1)
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	parts := strings.Fields(scanner.Text())
	if len(parts) < 2 {
		return
	}
	n, _ := strconv.Atoi(parts[0])
	k, _ := strconv.Atoi(parts[1])

	if !scanner.Scan() {
		return
	}
	numStr := strings.Fields(scanner.Text())
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i], _ = strconv.Atoi(numStr[i])
	}

	if backtrack(arr, k, 0) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
