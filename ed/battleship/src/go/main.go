package main

import (
	"bufio"
	"fmt"
	"os"
)

// Função que será chamada no LeetCode
func countBattleships(board [][]byte) int {
	totalNavios := 0
	totalLinhas := len(board)
	if totalLinhas == 0 {
		return 0
	}
	totalColunas := len(board[0])

	for linha := 0; linha < totalLinhas; linha++ {
		for coluna := 0; coluna < totalColunas; coluna++ {
			if board[linha][coluna] == '.' {
				continue
			}
			if linha > 0 && board[linha-1][coluna] == 'X' {
				continue
			}
			if coluna > 0 && board[linha][coluna-1] == 'X' {
				continue
			}
			totalNavios++
		}
	}

	return totalNavios

}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)

	board := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}

	result := countBattleships(board)
	fmt.Println(result)
}
