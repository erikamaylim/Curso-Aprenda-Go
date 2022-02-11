package main

import "fmt"

// Deslocamento de bits é quando deslocamos digitos binários para a esquerda ou direita.

func main() {
    
    a := 15 
    
    
	fmt.Printf("%v, %b\n", a, a)

	g := a >> 1
	fmt.Printf("%v, %b", g, g)
}
