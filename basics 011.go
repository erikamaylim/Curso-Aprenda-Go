
package main
import "fmt"

func main() {
    
    // var i uint16
    // i = 65536
    // fmt.Printf("%v", i) // dá overflow. uint16 vai até 65535.

    var i uint16
    i = 65535
    i++
    fmt.Printf("%v", i) // printa o número 0. Segue o mesmo princípio do odômetro. Após atingir o número máximo, retorna ao zero.
}