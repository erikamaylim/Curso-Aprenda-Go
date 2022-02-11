package main

import "fmt"

// usando switch com default e fallthrough:

func main() {
    sabores := []string{"chocolate", "baunilha", "morango", "banana"}

    for _, sabor := range sabores {
        switch sabor {
        case "morango":
            fmt.Println("Sorvete de",sabor, "é o meu favorito!")
            fallthrough
        case "baunilha", "chocolate":
            fmt.Println("Sorvete de",sabor, "é ótimo!")
        default:
            fmt.Println("Nunca experimentei o de", sabor, "antes.")
        }
    }
}