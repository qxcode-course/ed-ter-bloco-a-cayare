package main
import "fmt"
func main() {
    var nome string 
    var idade int

    fmt.Scan(&nome,&idade)

    if idade < 12{

        fmt.Println(nome, "eh crianca")
        
    } else if idade > 11 && idade < 18 {

        fmt.Println(nome, "eh jovem")
       
    } else if idade > 17 && idade < 65 {

        fmt.Println(nome, "eh adulto")
        
    } else if idade > 64 && idade < 1000{
        fmt.Println(nome, ("eh idoso"))
        
    } else {
         fmt.Println(nome, "eh mumia") 
         
    }
}

