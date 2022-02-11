
package main
import "fmt"

func main() {
    
    s := `Hello 
                crazy
                           world` // Para não dar erro em texto com formatação maluca, usar acento grave ao invés de aspas.
    fmt.Printf("%v\n%T\n", s, s)
    r := []byte(s)
    fmt.Printf("%v\n%T\n", r, r)
}