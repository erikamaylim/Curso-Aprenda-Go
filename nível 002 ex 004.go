/*  - Crie um programa que:
    - Atribua um valor int a uma variável
    - Demonstre este valor em decimal, binário e hexadecimal
    - Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
    - Demonstre esta outra variável em decimal, binário e hexadecimal */
package main
import "fmt"


func main() {
    x := 15
    fmt.Printf("%v em:\n - Decimal: %d\n - Binário: %b\n - Hexadecimal: %#X\n", x, x, x, x)
    
    fmt.Println("")
    
    y := x << 1
    fmt.Printf("%v em:\n - Decimal: %d\n - Binário: %b\n - Hexadecimal: %#X\n", y, y, y, y)
}