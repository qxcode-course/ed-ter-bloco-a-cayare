package main
import "fmt"
import "math"
func main() {
    var lado1 float64
    var lado2 float64
    var lado3 float64


    fmt.Scan(&lado1)
    fmt.Scan(&lado2)
    fmt.Scan(&lado3)

    r := (lado1+lado2+lado3) / 2

    area:= math.Sqrt(r* (r-lado1)* (r-lado2)* (r-lado3))
    fmt.Printf("%.2f\n",area)
}
