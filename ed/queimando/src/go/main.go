package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l int
	c int
}

func burnTrees(grid [][]rune, l, c int) {
	nl := len(grid)
	if nl == 0 {
		return
	}
	nc := len(grid[0])

	stack := NewStack[Pos]()
	stack.Push(Pos{l: l, c: c})

	for !stack.IsEmpty() {
		curr := stack.Pop()

		if curr.l < 0 || curr.l >= nl || curr.c < 0 || curr.c >= nc {
			continue
		}

		if grid[curr.l][curr.c] != '#' {
			continue
		}

		grid[curr.l][curr.c] = 'o'

		stack.Push(Pos{l: curr.l - 1, c: curr.c})
		stack.Push(Pos{l: curr.l + 1, c: curr.c})
		stack.Push(Pos{l: curr.l, c: curr.c - 1})
		stack.Push(Pos{l: curr.l, c: curr.c + 1})
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
