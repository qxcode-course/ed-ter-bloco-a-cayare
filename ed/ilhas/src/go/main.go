package main

import (
	"bufio"
	"fmt"
	"os"
)
func dfs (grid [][]byte, r int, c int){
	nr := len(grid)
	nc := len(grid[0])

	if r < 0 || c < 0 || r >= nr || c >= nc || grid[r][c] == '0'{
		return
	}
	grid[r][c] = '0'
	dfs(grid, r-1, c)
	dfs(grid, r+1, c)
	dfs(grid, r, c-1)
	dfs(grid, r, c+1)
}

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {
	if len(grid)== 0 {
		return 0
	}
	nr := len(grid)
	nc := len(grid[0])
	numIslands := 0

	for r := 0; r < nr; r++{
		for c := 0; c < nc; c++{
			if grid[r][c] == '1'{
				numIslands++
				dfs(grid, r, c)
			}
		}
	}
	return numIslands
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
