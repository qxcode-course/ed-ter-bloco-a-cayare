package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func jogadaValida(tabuleiro [][]rune, linha int, coluna int, numero rune) bool {
	tamanho := len(tabuleiro)

	for i := 0; i < tamanho; i++ {
		if tabuleiro[linha][i] == numero || tabuleiro[i][coluna] == numero {
			return false
		}
	}

	tamanhoQuadrante := 3
	if tamanho == 4 {
		tamanhoQuadrante = 2
	}

	linhaInicio := (linha / tamanhoQuadrante) * tamanhoQuadrante
	colunaInicio := (coluna / tamanhoQuadrante) * tamanhoQuadrante

	for l := 0; l < tamanhoQuadrante; l++ {
		for c := 0; c < tamanhoQuadrante; c++ {
			if tabuleiro[linhaInicio+l][colunaInicio+c] == numero {
				return false
			}
		}
	}

	return true
}

func resolverSudoku(tabuleiro [][]rune, linha int, coluna int) bool {
	tamanho := len(tabuleiro)

	if linha == tamanho {
		return true
	}

	if coluna == tamanho {
		return resolverSudoku(tabuleiro, linha+1, 0)
	}

	if tabuleiro[linha][coluna] != '.' {
		return resolverSudoku(tabuleiro, linha, coluna+1)
	}

	for numero := '1'; numero <= '0'+rune(tamanho); numero++ {
		if jogadaValida(tabuleiro, linha, coluna, numero) {
			tabuleiro[linha][coluna] = numero
			if resolverSudoku(tabuleiro, linha, coluna+1) {
				return true
			}
			tabuleiro[linha][coluna] = '.'
		}
	}
	return false
}

func main() {
	leitor := bufio.NewScanner(os.Stdin)
	if !leitor.Scan() {
		return
	}
	tamanhoStr := strings.TrimSpace(leitor.Text())
	tamanho, _ := strconv.Atoi(tamanhoStr)

	tabuleiro := make([][]rune, tamanho)
	for i := 0; i < tamanho; i++ {
		if !leitor.Scan() {
			return
		}
		linhaTexto := strings.TrimSpace(leitor.Text())
		tabuleiro[i] = []rune(linhaTexto)
	}

	if resolverSudoku(tabuleiro, 0, 0) {
		for i := 0; i < tamanho; i++ {
			fmt.Println(string(tabuleiro[i]))
		}
	}
}
