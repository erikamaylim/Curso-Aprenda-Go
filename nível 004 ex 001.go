/* - Usando uma literal composta:
     - Crie um array que suporte 5 valores to tipo int
     - Atribua valores aos seus índices 
     - Utilize range e demonstre os valores do array.
     - Utilizando format printing, demonstre o tipo do array.*/
package main
import "fmt"

func main() {
    meuArray := [5]int {}
    fmt.Println(meuArray)
    meuArray[0], meuArray[1], meuArray[2], meuArray[3], meuArray[4] = 1, 2, 3, 4, 5
    for _, v := range meuArray {
        fmt.Println(v)
    }
    fmt.Printf("%T", meuArray)
}
