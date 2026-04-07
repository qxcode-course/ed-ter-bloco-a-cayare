package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processa(vet []int) {
	// 1. PONTO DE PARADA:
	if len(vet) < 1 {
		return
	}

	// 2. MONTE O VETOR AUXILIAR

	if len(vet) > 1 {
		var aux []int
		// O limite 'len(vet) - 1' é CRUCIAL para não tentar somar o vento
		for i := 0; i < len(vet)-1; i++ {
			soma := vet[i] + vet[i+1]
			aux = append(aux, soma)
		}

		// 3. CHAMADA RECURSIVA:
		processa(aux)
	}

	// 4. IMPRESSÃO:

	fmt.Printf("[ %s ]\n", Join(vet, " "))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	line := scanner.Text()
	parts := strings.Fields(line)
	vet := []int{}
	for _, part := range parts {
		if value, err := strconv.Atoi(part); err == nil {
			vet = append(vet, value)
		}
	}
	processa(vet)
}

func Join[T any](v []T, sep string) string {
	if len(v) == 0 {
		return ""
	}
	s := ""
	for i, x := range v {
		if i > 0 {
			s += sep
		}
		s += fmt.Sprintf("%v", x)
	}
	return s
}
