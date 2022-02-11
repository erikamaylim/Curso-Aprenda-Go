package main
import "fmt"

// Golang não tem while
// Mais sobre o uso de for: https://go.dev/doc/effective_go#for

func main() {
    x := 0
    for {
        if x < 10 {
            fmt.Println("O valor de x é", x ,"e ele é menor que 10")
            x ++
        } else {
            fmt.Println("x é igual a 10")
            break
        }
    }
} 
