
package main
import "fmt"

func main() {
    
    a := "e"
    b := "é"
    c := "☼" 
    
    fmt.Printf("%v, %v, %v\n", a, b, c)
    
    d := []byte(a) // convertendo string em byte slice
    e := []byte(b)
    f := []byte(c)
    
        fmt.Printf("%v, %v, %v", d, e, f)

    
}