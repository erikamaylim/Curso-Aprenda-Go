/*  - Crie uma variável de tipo string utilizando uma raw string literal.
    - Demonstre-a. */
package main
import "fmt"


func main() {
    
    s := `This is a\nRaw String Literal\n. See? \n doesn't work here.
    
                Weird                   
                           as                   
                                   f*ck!`
    fmt.Print(s)
}