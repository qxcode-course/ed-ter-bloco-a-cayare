package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func backtrack(chars []rune, current []rune, visited []bool) {
	if len(current) == len(chars) {
		fmt.Println(string(current))
		return
	}
	for i := 0; i < len(chars); i++ {
		if visited[i] {
			continue
		}

		visited[i] = true
		current = append(current, chars[i])
		backtrack(chars, current, visited)

		current = current[:len(current)-1]
		visited[i] = false
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	s := strings.TrimSpace(scanner.Text())

	chars := []rune(s)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	visited := make([]bool, len(chars))
	var current []rune

	backtrack(chars, current, visited)

}
