/*  - Crie um loop utilizando a sintaxe: for {}
    - Utilize-o para demonstrar os anos desde que você nasceu.*/
package main

import (
	"fmt"
)

func main() {
	anonasc := 1989
	anoatual := 2021
	for {
	    if anonasc > anoatual{
	        break
	    }
    fmt.Println(anonasc)
		anonasc++
	}
	
}