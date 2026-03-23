package main

import "fmt"

type jogada struct {
	pa, pb int
}

func abs(j jogada) int {
	if j.pa > j.pb {
		return j.pa - j.pb
	}
	return j.pb - j.pa
}

func main() {
	qtd := 0
	fmt.Scan(&qtd)
	jogadas := make([]jogada, qtd)
	for i := range jogadas {
		fmt.Scan(&jogadas[i].pa, &jogadas[i].pb)

	}
	indMelhor := -1

	for i, jog := range jogadas {
		if jog.pa < 10 || jog.pb < 10 {
			continue
		}
		if indMelhor == -1 {
			indMelhor = i
		} else if abs(jog) < abs(jogadas[indMelhor]) {
			indMelhor = i
		}
	}
	if indMelhor == -1 {
		fmt.Println("sem ganhador")
	} else {
		fmt.Println(indMelhor)
	}
}
