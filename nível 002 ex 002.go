/* Escreva expressões utilizando os seguintes operadores:
        - ==
        - <
        - >
        - <=
        - >=
        - !=
   e atribua seus valores a variáveis.
   Demonstre estes valores. */
package main
import "fmt"

func main() {
    
    x := 8
    y := 12
    a := x == y
    b := x < y
    c := x > y
    d := x <= y
    e := x >= y
    f := x != y
    
    fmt.Printf("a)%v = %v? %v\n", x, y, a)
    fmt.Printf("b)%v < %v? %v\n", x, y, b)
    fmt.Printf("c)%v > %v? %v\n", x, y, c)
    fmt.Printf("d)%v <= %v? %v\n", x, y, d)
    fmt.Printf("e)%v >= %v? %v\n", x, y, e)
    fmt.Printf("f)%v != %v? %v\n", x, y, f)
    
    
}