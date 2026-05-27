package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isValid(arr []byte, index int, val byte, l int) bool {
	for i := index - l; i <= index+l; i++ {
		if i >= 0 && i < len(arr) && i != index {
			if arr[i] == val {
				return false
			}
		}
	}
	return true
}
func backtrack(arr []byte, index int, l int) bool {
	if index == len(arr) {
		return true
	}
	if arr[index] != '.' {
		return backtrack(arr, index+1, l)
	}
	for d := 0; d <= l; d++ {
		val := byte('0' + d)
		if isValid(arr, index, val, l) {
			arr[index] = val
			if backtrack(arr, index+1, l) {
				return true
			}
			arr[index] = '.'
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	seq := strings.TrimSpace(scanner.Text())
	if !scanner.Scan() {
		return
	}
	lStr := strings.TrimSpace(scanner.Text())
	l, _ := strconv.Atoi(lStr)

	arr := []byte(seq)
	if backtrack(arr, 0, l) {
		fmt.Println(string(arr))
	}
}
