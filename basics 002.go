package main
import "fmt"

var y = 10
func main() {
    
    z := 20
    qualquercoisa(z)
    fmt.Println("z =", z)
}

func qualquercoisa(x int) {
    fmt.Println("x =", x )
    fmt.Println("y =", y)
}