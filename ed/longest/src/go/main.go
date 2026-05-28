package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var direcoes = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func buscarMaiorCaminho(matriz [][]int, blocoDeNotas [][]int, linha int, coluna int) int {
	if blocoDeNotas[linha][coluna] != 0 {
		return blocoDeNotas[linha][coluna]
	}

	tamanhoMaximo := 1

	for _, dir := range direcoes {
		proximaLinha := linha + dir[0]
		proximaColuna := coluna + dir[1]

		if proximaLinha >= 0 && proximaColuna >= 0 && proximaLinha < len(matriz) && proximaColuna < len(matriz[0]) {
			if matriz[proximaLinha][proximaColuna] > matriz[linha][coluna] {
				comprimentoDoCaminho := 1 + buscarMaiorCaminho(matriz, blocoDeNotas, proximaLinha, proximaColuna)
				if comprimentoDoCaminho > tamanhoMaximo {
					tamanhoMaximo = comprimentoDoCaminho
				}
			}
		}
	}
	blocoDeNotas[linha][coluna] = tamanhoMaximo
	return tamanhoMaximo
}

func encontrarCaminhoCrescenteMaisLongo(matriz [][]int) int {
	if len(matriz) == 0 || len(matriz[0]) == 0 {
		return 0
	}

	totalLinhas := len(matriz)
	totalColunas := len(matriz[0])

	blocoDeNotas := make([][]int, totalLinhas)
	for i := range blocoDeNotas {
		blocoDeNotas[i] = make([]int, totalColunas)
	}

	maiorCaminhoGeral := 0

	for i := 0; i < totalLinhas; i++ {
		for j := 0; j < totalColunas; j++ {
			caminhoAtual := buscarMaiorCaminho(matriz, blocoDeNotas, i, j)
			if caminhoAtual > maiorCaminhoGeral {
				maiorCaminhoGeral = caminhoAtual
			}
		}
	}
	return maiorCaminhoGeral
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	parts := strings.Fields(scanner.Text())
	if len(parts) < 2 {
		return
	}
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	matrix := make([][]int, nl)
	for i := 0; i < nl; i++ {
		if !scanner.Scan() {
			return
		}
		tokens := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc && j < len(tokens); j++ {
			v, _ := strconv.Atoi(tokens[j])
			row[j] = v
		}
		matrix[i] = row
	}

	fmt.Println(encontrarCaminhoCrescenteMaisLongo(matrix))
}
