/*  - Demonstre o resto da divisão por 4 de todos os números entre 10 e 100 */
package main

import (
	"fmt"
)

func main() {
    
    for x := 10; x <= 100; x++ {
        fmt.Println("Na divisão entre", x, "e 4, o resto da divisão é", x % 4)
    }
}