package main

import "fmt"

func main() {

	fila := NewQueue[string]()

	for c := 'A'; c <= 'P'; c++ {

		fila.Enqueue(string(c))

	}

	for fila.items.Len() > 1 {
		time1 := fila.Dequeue()
		time2 := fila.Dequeue()

		var gols1, gols2 int
		fmt.Scan(&gols1, &gols2)

		if gols1 > gols2 {
			fila.Enqueue(time1)
		} else {
			fila.Enqueue(time2)
		}
	}

	campeao := fila.Dequeue()
	fmt.Println(campeao)

}
