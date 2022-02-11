/* - Desafio surpresa!
- Format printing:
    - Decimal       %d
    - Hexadecimal   %#x
    - Unicode       %#U ou %c (exibe apenas o caractere)
    - Tab           \t
    - Linha nova    \n
- Faça um loop dos números 33 a 122, e utilize format printing para demonstrá-los como texto/string. */

package main
import "fmt"

func main() {
    for x := 33; x <= 122; x ++ {
        fmt.Printf("Nº %v em:\n\t\tDecimal: %d\tHexadecimal: %#X\tUnicode: %c\n", x, x, x, x)
        fmt.Println("___________________________________________________________________________________" )
    }
}    