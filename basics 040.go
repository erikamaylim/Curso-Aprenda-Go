/* slice: make */
package main
import "fmt"

func main() {
    slice := make([]int, 5, 10)
    fmt.Println(slice)
    slice[0], slice [1], slice[2], slice[3] = 1, 2, 3, 4
    fmt.Println(slice)
    slice[4] = 5
    fmt.Println(slice,"\nTamanho (len):", len(slice), "\nCapacidade (cap):", cap(slice))
    slice = append(slice, 6, 7, 8, 9, 10) // x[n] onde n é maior que len é out of range. Use append.
    fmt.Println(slice,"\nTamanho (len):", len(slice), "\nCapacidade (cap):", cap(slice))
    slice = append(slice, 11) // Append maior que cap modifica o array subjacente.
    fmt.Println(slice,"\nTamanho (len):", len(slice), "\nCapacidade (cap):", cap(slice))
}