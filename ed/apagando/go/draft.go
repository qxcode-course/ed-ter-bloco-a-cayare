package main

import "fmt"

func main() {

	n := 0
	fmt.Scan(&n)
	vetor := make([]int, n)
	for i := 0; i < n; i++ {

		fmt.Scan(&vetor[i])
	}

	m := 0
	fmt.Scan(&m)
	apagados := make(map[int]bool)
	for i := 0; i < m; i++ {
		var x int
		fmt.Scan(&x)
		apagados[x] = true

	}
	for i := 0; i < n; i++ {
		if !apagados[vetor[i]] {
			fmt.Print(vetor[i], " ")
		}
	}
	fmt.Println()
}
