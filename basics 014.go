package main
import "fmt"

func main() {
    
    s := "Hello World éàüâ" 
    
    for _, v := range s { // desta forma, exibe por caractere.
        fmt.Printf("%v   |   %T   |   %#U   |   %#x\n", v, v, v, v)
    }    
        
    fmt.Println("")
    
    for i := 0; i < len(s); i++ { // desta forma, exibe por byte. 
        fmt.Printf("%v   |   %T   |   %#U   |   %#x\n", s[i], s[i], s[i], s[i])
    }
}