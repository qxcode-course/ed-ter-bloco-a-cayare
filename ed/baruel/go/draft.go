package main

import "fmt"

func repetidos(vetor []int) {

	achou := false

	for i := 0; i < len(vetor); i++ {

		cont := 0

		for j := i + 1; j < len(vetor); j++ {

			if vetor[i] == vetor[j] {
				cont++
			}
		}

		if cont > 0 {

			for k := 0; k < cont; k++ {

				if achou {
					fmt.Print(" ")
				}

				fmt.Print(vetor[i])
				achou = true
			}

			i += cont
		}
	}

	if !achou {
		fmt.Println("N")
	} else {
		fmt.Println()
	}
}

func falta(x int, vetor []int) {

	achou := false

	for num := 1; num <= x; num++ {

		existe := false

		for i := 0; i < len(vetor); i++ {

			if num == vetor[i] {
				existe = true
				break
			}
		}

		if !existe {

			if achou {
				fmt.Print(" ")
			}

			fmt.Print(num)
			achou = true
		}
	}

	if !achou {
		fmt.Println("N")
	} else {
		fmt.Println()
	}

}

func main() {

	var figN int
	var figB int

	fmt.Scan(&figN, &figB)

	vetor := make([]int, figB)

	for i := 0; i < figB; i++ {
		fmt.Scan(&vetor[i])
	}

	repetidos(vetor)
	falta(figN, vetor)

}
