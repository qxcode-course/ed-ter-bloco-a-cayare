package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	r, c int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	line := scanner.Text()
	var rows, cols int
	fmt.Sscanf(line, "%d %d", &rows, &cols)

	matrix := make([][]rune, rows)
	var start, end Point

	for i := 0; i < rows; i++ {
		if !scanner.Scan() {
			break
		}
		l := scanner.Text()
		matrix[i] = []rune(l)
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 'I' {
				start = Point{i, j}
			} else if matrix[i][j] == 'F' {
				end = Point{i, j}
			}
		}
	}

	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}
	caminho := NewStack[Point]()
	becos := NewStack[Point]()
	caminho.Push(start)

	dr := []int{-1, 1, 0, 0}
	dc := []int{0, 0, -1, 1}

	for !caminho.IsEmpty() {
		atual := caminho.Top()
		visited[atual.r][atual.c] = true
		matrix[atual.r][atual.c] = '.'

		if atual == end {
			break
		}

		var validos []Point
		for i := 0; i < 4; i++ {
			nr := atual.r + dr[i]
			nc := atual.c + dc[i]

			if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
				if matrix[nr][nc] != '#' && !visited[nr][nc] {
					validos = append(validos, Point{nr, nc})
				}
			}
		}
		if len(validos) > 0 {
			caminho.Push(validos[0])
		} else {
			becos.Push(atual)
			caminho.Pop()
		}
	}
	for !becos.IsEmpty() {
		b := becos.Pop()
		matrix[b.r][b.c] = ' '
	}

	for i := 0; i < rows; i++ {
		fmt.Println(string(matrix[i]))
	}
}
