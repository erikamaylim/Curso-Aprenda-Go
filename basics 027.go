package main
import "fmt"

// usando switch:
func main() {
    x := 11
    
    switch {
        case x < 10:
            fmt.Println("x é menor que 10")
        case x == 10:
            fmt.Println("x é igual a 10")
        case x > 10:
            fmt.Println("x é maior que 10")

    }
}