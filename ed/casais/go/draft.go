package main

import "fmt"

func main() {
	var quantidadeAnimal int

	fmt.Scan(&quantidadeAnimal)
	animal := make([]int, quantidadeAnimal)

	for i := 0; i < quantidadeAnimal; i++ {
		fmt.Scan(&animal[i])

	}
	cont := 0
	for i := 0; i < quantidadeAnimal; i++ {

		if animal[i] == 0 {
			continue
		}

		for j := 0; j < quantidadeAnimal; j++ {

			if animal[i] == animal[j]*-1 {
				cont++
				animal[j] = 0
				animal[i] = 0
				break
			}
		}
	}
	fmt.Println(cont)
}
