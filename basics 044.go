/* MAPS: range & deletando */
package main
import "fmt"

func main() {
    exemplo := map[int] string {
        1: "primeiro",
        2: "segundo",
        3: "terceiro",
        4: "quarto",
    }
    fmt.Println(exemplo)
    for key, value := range exemplo {
        fmt.Println(key, value)
    }
    
    delete(exemplo, 1) 
    
    fmt.Println(exemplo)
    for key, value := range exemplo {
        fmt.Println(key, value)
    }    
}

/* Observações:
    - Range: for k, v := range map { }
    - Reiterando: maps *não tem ordem* e um range usará uma ordem aleatória.*/