/* Escreva um programa que mostre um número em decimal, binário, e hexadecimal. */
package main
import "fmt"

func main() {
    x := 106
    fmt.Printf("x em decimal: %d\nx em binário: %b\nx em hexadecimal: %#X", x, x, x)
}