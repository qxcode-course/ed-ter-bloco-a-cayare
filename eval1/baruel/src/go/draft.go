package main

import "fmt"



func main() {
	var album int
	var figB int
	fmt.Scan(&album)
	fmt.Scan(&figB)
    fig := make([]int, figB)

	for i := 0; i < figB; i++ {
		fmt.Scan(&fig[i])
	}
    repetido := 0
    for i:=0; i < figB - 1; i++{

        if fig[i] != fig[i+1]{
            fig[i] = 0
        } 

        if fig[i] != 0 {
            fmt.Print(fig[i], " ")
            repetido++
        }
    }
    fmt.Printf("\n")
    
}
