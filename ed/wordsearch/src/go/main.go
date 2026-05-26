package main

import (
	"bufio"
	"fmt"
	"os"
)

func dfs(grid [][]byte, word string, r int, c int, index int) bool {
	if index == len(word) {
		return true
	}

	if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || grid[r][c] != word[index] {
		return false
	}
	temp := grid[r][c]
	grid[r][c] = '#'

	found := dfs(grid, word, r+1, c, index+1) ||
		dfs(grid, word, r-1, c, index+1) ||
		dfs(grid, word, r, c+1, index+1) ||
		dfs(grid, word, r, c-1, index+1)

	grid[r][c] = temp

	return found
}

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	if len(grid) == 0 || len(word) == 0 {
		return false
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == word[0] {
				if dfs(grid, word, i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
