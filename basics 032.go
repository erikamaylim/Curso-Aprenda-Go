package main
import "fmt"

// Operadores lógicos condicionais

func main() {
    
    x := 4
    if (x % 2 == 0 || x % 3 == 0) {
        fmt.Println("x é múltiplo de 2 ou múltiplo de 3")
    }
    
    y := 6
    if (y % 2 == 0 && y % 3 == 0) {
        fmt.Println("y é múltiplo de 2 e também de 3")
    }
}