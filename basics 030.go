package main

import "fmt"

// switch com cases compostos:

func main() {
    
    estado := "Texas"
    
    switch estado {
        case "Amazonas", "Roraima", "Amapá", "Pará", "Tocantins", "Rondônia", "Acre":
            fmt.Println(estado, "fica no Norte do Brasil")
        case "Maranhão", "Piauí", "Ceará", "Rio Grande do Norte", "Pernambuco", "Paraíba", "Sergipe", "Alagoas", "Bahia":
            fmt.Println(estado, "fica no Nordeste do Brasil")
        case "Distrito Federal":
            fmt.Println(estado,"não é um Estado.")
            fallthrough
        case "Mato Grosso", "Mato Grosso do Sul", "Goiás":
            fmt.Println(estado, "fica no Centro-Oeste do Brasil")
        case "São Paulo", "Rio de Janeiro", "Espírito Santo", "Minas Gerais":
            fmt.Println(estado, "fica no Sudeste do Brasil")
        case "Paraná", "Rio Grande do Sul", "Santa Catarina":
            fmt.Println(estado, "fica no Sul do Brasil")
        default:
            fmt.Println(estado,"não é um estado brasileiro")
    }    
}
