/*  - Crie um programa que utilize a declaração switch, sem switch statement especificado. */
package main

import (
	"fmt"
)

func main() {
    
    x := 1
    switch {
        case x == 1:
            fmt.Println("Um maluco no pedaço")
        case x == 2:
            fmt.Println("Família Dinossauro")
        case x == 3:
            fmt.Println("Beavis and Butt-Head")
        case x == 4:
            fmt.Println("Futurama")
        default:
            fmt.Println("Vá ler um livro!")
    }
}