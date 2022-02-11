/*  - Crie um programa que demonstre o funcionamento da declaração if. */
package main

import (
	"fmt"
)

func main() {
    
    for x := 20; x <= 100; x++ {
        if x % 7 == 0 {
            fmt.Println(x, "é divisível por 7")
        } 
    }
}