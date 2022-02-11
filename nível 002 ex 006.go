/*  - Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
    - Demonstre estes valores. */
package main
import "fmt"

const (
    _ = iota  
    a = 2021 + iota
    b 
    c 
    d 
)

func main() {
    
    fmt.Printf("%v\n%v\n%v\n%v", a, b, c, d)
}