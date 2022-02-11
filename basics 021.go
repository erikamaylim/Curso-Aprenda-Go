package main
import "fmt"

// Loops: nested loop (repetição hierárquica)

func main() {
    for hora := 0; hora < 12; hora ++ {
        fmt.Printf("Hora: %v\n", hora)
        for x := 0; x < 60; x ++ { 
            fmt.Print(x, " ")
        }
        fmt.Println(`
        `)
    }    
} 

