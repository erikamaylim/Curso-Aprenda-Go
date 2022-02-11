
package main
import "fmt"

// O tipo de uma constante só é definido no momento do seu uso.
// O tipo de uma variável é definido no momento da atribuição.

const x = 10
var y = 20
var c int
var d float64

func main() {
    
    //c = x                               // Resultado: 10, int
    //fmt.Printf("%v, %T", c, c)
    //d = x                              // Resultado: 10, float64
    //fmt.Printf("%v, %T", d, d)
    
    //c = y                                 // Resultado: 20, int
    //fmt.Printf("%v, %T", c, c)
    //d = y                              // dá erro, pois y é int e d é float
    //fmt.Printf("%v, %T", d, d)
}