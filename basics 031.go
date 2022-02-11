package main
import "fmt"

var x interface{}

func main() {
    
    x = false
    switch x.(type) {
        case int:
            fmt.Println("x é do tipo int")
        case bool:
            fmt.Println("x é do tipo bool")
        case string:
            fmt.Println("x é do tipo string")
        case float64:
            fmt.Println("x é do tipo float")
        
    }
    
}