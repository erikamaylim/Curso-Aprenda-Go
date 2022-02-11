package main

import "fmt"

// criando seu próprio tipo

type hotdog int
var b hotdog = 10

func main() {
    x := 20
    fmt.Printf("%v, %T\n", x, x)
    fmt.Printf("%v, %T", b, b)

}
