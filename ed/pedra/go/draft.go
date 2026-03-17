package main

import "fmt"

func main() {
    var n int 
    fmt.Scan(&n)

    var ganhador_indice int = -1
    var menor_diferenca int = 999

    for i := 0; i < n; i++{
        var a, b int 
        fmt.Scan(&a, &b)

        if a >= 10 && b >= 10 {
            dif := a - b
           
            if dif < 0 {
                dif = - dif
            } 

            if dif < menor_diferenca {
                menor_diferenca = dif
                ganhador_indice = i
            }

        } 

    }   
if ganhador_indice != -1{
    
    fmt.Println(ganhador_indice)
} else {
    fmt.Println("sem ganhador")
}




    
}

