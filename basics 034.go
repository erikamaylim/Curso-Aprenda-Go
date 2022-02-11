/* SLICES 
    - Slice é uma estrutura que aponta para um array, permitindo manipular os seus elementos e o seu tamanho. Na
        verdade, um slice não armazena nenhum valor: sempre haverá um array associado onde os elementos estarão
        armazenados.
    - É possível manipular os elementos do slice da mesma forma que nos arrays (através dos índices), 
        com a vantagem de que slices possuem funções para adicionar elementos, alterando seus tamanhos */

package main
import "fmt"

func main() {

    slice := []int {1, 2, 3, 4, 5}
    fmt.Println("slice =", slice)
    slice2 := append(slice, 6) // Não é possível fazer isso com um array
    fmt.Println("slice2 =", slice2)
    
    slice[3] = 520 // É possível alterar o valor dos elementos
    fmt.Println("slice =", slice)
    

}