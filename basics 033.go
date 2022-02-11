
/* ARRAYS 
    - Armazena variáveis do mesmo tipo
    - Cada elemento é identificado por sua posição dentro do array
    - A quantidade de elementos (tamanho do array) e seu tipo são definidas no momento da declaração
    - São imutáveis
    - Os índices começam em 0 (zero) e vão até (tamanho-1)*/

package main
import "fmt"

var x [5]int

func main() {
    x[0] = 1
    x[1] = 10
    fmt.Println(x[0])
    fmt.Println(x[1])
    fmt.Println(x[2])
    fmt.Println(x)
    fmt.Println(len(x))
    fmt.Printf("%T", x)
}