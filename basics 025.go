package main
import "fmt"

// formas de usar o if:
func main() {
    x := 10
    if x < 100 {
        fmt.Println("x é menor que 100")
    } 
    
    y := 20
    if !(y < 100){
        fmt.Println("y é menor que 100")
    } 
    
    if z:= 30; z < 100 {
        fmt.Println("z é menor que 100")
    }
}