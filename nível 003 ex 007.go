/*  - Utilizando a solução anterior, adicione as opções else if e else. */
package main

import (
	"fmt"
)

func main() {
    
    for x := 20; x <= 100; x++ {
        if x % 7 == 0 && x % 9 == 0 {
            fmt.Println(x, "é divisível por 7 e por 9!!!!")
        } else if x % 7 == 0 {
            fmt.Println(x, "é divisível apenas por 7")
        } else if x % 9 == 0 {
            fmt.Println(x, "é divisível apenas por 9")
        } else {
            fmt.Println(x, "NÃO é divisível por 7 e nem por 9")
        }
    }
}