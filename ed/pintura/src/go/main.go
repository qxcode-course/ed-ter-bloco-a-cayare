package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func dfs(image [][]int, r int, c int, originalColor int, newColor int) {
	nr := len(image)
	nc := len(image[0])

	if r < 0 || c < 0 || r >= nr || c >= nc || image[r][c] != originalColor {
		return
	}

	image[r][c] = newColor

	dfs(image, r-1, c, originalColor, newColor)
	dfs(image, r+1, c, originalColor, newColor)
	dfs(image, r, c-1, originalColor, newColor)
	dfs(image, r, c+1, originalColor, newColor)
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	originalColor := image[sr][sc]
	if originalColor != color {
		dfs(image, sr, sc, originalColor, color)
	}
	return image
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	line := scanner.Text()
	parts := strings.Fields(line)
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	image := make([][]int, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		rowStr := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc; j++ {
			row[j], _ = strconv.Atoi(rowStr[j])
		}
		image[i] = row
	}

	if scanner.Scan() {
		params := strings.Fields(scanner.Text())
		sr, _ := strconv.Atoi(params[0])
		sc, _ := strconv.Atoi(params[1])
		color, _ := strconv.Atoi(params[2])

		result := floodFill(image, sr, sc, color)

		for _, row := range result {
			for j, val := range row {
				if j > 0 {
					fmt.Print(" ")
				}
				fmt.Print(val)
			}
			fmt.Println()
		}
	}
}