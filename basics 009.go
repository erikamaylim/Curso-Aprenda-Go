package main

import "fmt"

// conversão de tipos

type hotdog int
var b hotdog = 202

func main() {
    fmt.Printf("b = %v, %T\n", b, b)
    x := 10
    fmt.Printf("x = %v, %T\n", x, x)
    
    x = int(b) // para x receber o valor de b é necessário converter o tipo de b para o mesmo de x
    fmt.Printf("x = %v, %T", x, x)
   
}
