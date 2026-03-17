package main
import "fmt"
func main() {
    var h int
    var p int
    var l int
    var d int
    var i int

    fmt.Scan(&h, &p, &l, &d)

    i = l

    for {

        if d == -1 {
            i--
            if i == -1 {
                i = 15
            }
        } else {
            i++
            if i == 16 {
                i = 0
            }
        }

        l = i

        if l == h {
            fmt.Println("S")
            break

        } else if l == p {
            fmt.Println("N")
            break
        }
    }
}