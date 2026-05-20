package main

import (
	"bufio"
	"fmt"
	"os"
)
func dfs(board [][]byte, r int, c int){
	nr := len(board)
	nc := len(board[0])

	if r < 0 || c < 0 || r >= nr || c >= nc || board[r][c] != 'O'{
		return
	}

	board[r][c] = '*'

	dfs(board, r-1, c)
	dfs(board, r+1, c)
	dfs(board, r, c-1)
	dfs(board, r, c+1)
}
// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	if len (board) == 0{
		return
	}

	nr := len(board)
	nc := len(board[0])

	for i := 0; i < nr; i++{
		if board[i][0] == 'O' {
			dfs(board, i, 0)
		}
		if board[i][nc-1] == 'O'{
			dfs(board, i, nc-1)
		}

	}
	for j := 0; j < nc; j++{
		if board[0][j] == 'O'{
			dfs(board, 0 , j)
		}
		if board[nr-1][j] == 'O'{
			dfs(board, nr-1, j)
		}

	}
	for i := 0; i < nr; i++{
		for j := 0; j < nc; j++{
			if board[i][j] == 'O'{
				board[i][j] = 'X'
			} else if board[i][j] == '*'{
				board[i][j] = 'O'
			}
		}
	}
}


// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
