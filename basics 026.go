package main
import "fmt"

// formas de usar o if - else / if - else if - else:
func main() {
    x := 10000
    if x < 100 {
        fmt.Println("x é menor que 100")
    } else {
        fmt.Println("x é maior que 100")
    }
    
    if y := 100; y < 100 {
        fmt.Println("y é menor que 100")
    } else if y > 100 {
        fmt.Println("y é maior que 100")
    } else {
        fmt.Println("y é igual a 100")
    }
}