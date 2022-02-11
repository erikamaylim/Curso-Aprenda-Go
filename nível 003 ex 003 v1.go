/*  - Crie um loop utilizando a sintaxe: for condition {}
    - Utilize-o para demonstrar os anos desde que você nasceu.*/
package main

import (
	"fmt"
)

func main() {
	anonasc := 1989
	anoatual := 2021
	for anonasc <= anoatual {
		fmt.Println(anonasc)
		anonasc++
	}
}