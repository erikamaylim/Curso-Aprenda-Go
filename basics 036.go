/* SLICES (cont2)
    - Range em slices */

package main
import "fmt"

func main() {

    slice := []int {64, 63, 35, 34, 32, 10}
    soma := 0
    for _, valor := range slice {
        soma += valor // é o mesmo que soma = soma + valor
    } 
    fmt.Println("A soma dos valores do slice é", soma)
    
}