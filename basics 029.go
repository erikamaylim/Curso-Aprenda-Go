package main

import "fmt"

// Outro exemplo de switch usando default:

func main() {
    animal := "babirusa"
    
    switch animal {
        case "cachorro":
            fmt.Println("Este bicho é um cachorro")
        case "gato":
            fmt.Println("Este bicho é um gato")
        case "papagaio":
            fmt.Println("Este bicho é um papagaio")
        default:
            fmt.Println("Eu não sei que bicho é esse")
    }    
}