/* APPEND */
package main
import "fmt"

func main() {
    
    slice1 := []int {1, 2, 3, 4, 5}
    slice2 := []int {9, 10, 11, 12}
    fmt.Println("slice1:", slice1)
    fmt.Println("slice2:", slice2)
    slice1 = append(slice1, 6, 7, 8) // acrescentando elementos à slice
    fmt.Println("slice1 após append:", slice1)
    slice1 = append(slice1, slice2...) /* Utiliza-se ... para utilizar os elementos da slice como variáveis individuais,
                                            pois o append não aceita slice como variável. */
    fmt.Println("slice1 após append com slice2:", slice1)
    
    
}