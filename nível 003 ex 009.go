/* - Crie um programa que utilize a declaração switch, 
        onde o switch statement seja uma variável do tipo string 
        com identificador "esporteFavorito". */
package main
import "fmt"


func main() {
    
    esporteFavorito := "j"
    switch esporteFavorito {
        case "B","b":
            fmt.Println("Eu jogo Basquete")
        case "F", "f":
            fmt.Println("Eu jogo Futebol")
        case "J", "j":
            fmt.Println("Eu luto Jiu-Jitsu")
        case "N", "n":
            fmt.Println("Eu faço Natação")
        default:
            fmt.Println("Eu sou sedentário")
    }
    
}