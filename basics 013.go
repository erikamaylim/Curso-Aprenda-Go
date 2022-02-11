
package main
import "fmt"

func main() {
    
    s := "Hello World" 
    r := []byte(s)
    for _, v := range r { 
        fmt.Printf("%v   |   %T   |   %#U   |   %#x\n", v, v, v, v)

    }
}