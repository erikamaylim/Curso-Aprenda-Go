/*  - Usando uma literal composta:
        - Crie uma slice de tipo int
        - Atribua 10 valores a ela
    - Utilize range para demonstrar todos estes valores.
    - E utilize format printing para demonstrar seu tipo. */
package main
import "fmt"

func main() {

    meuSlice := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
    fmt.Println(meuSlice)
    for i, v := range meuSlice {
        fmt.Println("índice", i,"-",v)
    }
    fmt.Printf("%T", meuSlice)
}
