package main
import "fmt"
func main() {
	pessoas := 0
	fmt.Scan(&pessoas)
	identificador := make([]int, pessoas)
	for i := 0; i < pessoas; i++ {
		fmt.Scan(&identificador[i])
	}
	desistentes := 0
	fmt.Scan(&desistentes)
	Vdesistentes := make([]int, desistentes)
	for i := 0; i < desistentes; i++ {
		fmt.Scan(&Vdesistentes[i])
	}
	for i := 0; i <pessoas; i++ {
		for j := 0; j < desistentes; j++ {
			if identificador[i] == Vdesistentes[j] {
				identificador[i] = 0
			}
		}
	}
	for i := 0; i < pessoas; i++ {
		if identificador[i] != 0 {
			fmt.Print(identificador[i], " ")
		}
	}
	fmt.Printf("\n")
}