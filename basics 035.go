/* SLICES (cont)
    - Range em slices */

package main
import "fmt"

func main() {

    slice := []int {1958, 1959, 1986, 1987, 1989}
    for índice, valor := range slice {
        fmt.Println("No índice", índice, "o valor é", valor)
            
    }    
    
    slice = append(slice, 2011)
    fmt.Println(slice)
    
    fmt.Print("Os anos armazenados no slice são: ")
    for _, valor := range slice {
        fmt.Printf("%v, ", valor)
    }
    
    

}