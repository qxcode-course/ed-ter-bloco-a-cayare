package main

import "fmt"

func imprimir_fila(fila []int, posEspada int) {
	fmt.Print("[ ")

	for i, valor := range fila {

		if i == posEspada {
			fmt.Printf("%d> ", valor)
		} else {
			fmt.Printf("%d ", valor)
		}
	}
	fmt.Print("]")
}
func main() {

	var n int
	var e int
	fmt.Scan(&n, &e)

	fila := make([]int, n)
	for i := 0; i < n; i++ {
		fila[i] = i + 1
	}
	espada := 0
	for i := 0; i < n; i++ {
		if fila[i] == e {
			espada = i
			break
		}

	}
	for len(fila) > 1 {
		imprimir_fila(fila, espada)

		vitima := (espada + 1) % len(fila)
		fila = append(fila[:vitima], fila[vitima+1:]...)

		espada = vitima % len(fila)
		fmt.Println()
	}
	imprimir_fila(fila, espada)
	fmt.Println()
}
