/*  - Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
        - Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
        - Do quinto ao último item do slice (incluindo o último item!)
        - Do segundo ao sétimo item do slice (incluindo o sétimo item!)
        - Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
    - Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item */
package main
import "fmt"

func main() {

    meuSlice := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
    fmt.Println(meuSlice)
 
    fmt.Println("1º ao 3º:", meuSlice[:3])
    fmt.Println("5º ao último:", meuSlice[4:])
    fmt.Println("2º ao 7º:", meuSlice[1:7])
    fmt.Println("3º ao penúltimo:", meuSlice[2:9])
    fmt.Println("3º ao penúltimo:", meuSlice[2:len(meuSlice) - 1])
}
