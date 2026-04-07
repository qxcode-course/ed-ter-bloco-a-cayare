package main

import "fmt"

func imprimir(n int) {
	if n == 0 {
		return
	}
	proximo := n / 2
	resto := n % 2
	imprimir(proximo)
	fmt.Printf("%d %d\n", proximo, resto)

}

func main() {
	var n1 int

	fmt.Scan(&n1)

	imprimir(n1)

}
