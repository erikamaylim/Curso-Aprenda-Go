/*  - Crie uma slice usando make que possa conter todos os estados do Brasil.
        - Os estados: "Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás",
        "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", 
        "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", 
        "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"
    - Demonstre o len e cap da slice.
    - Demonstre todos os valores da slice *sem utilizar range.* */
package main
import "fmt"

func main() {

    Brasil := make([]string, 26, 26)
    Brasil = []string {"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás",
        "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", 
        "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", 
        "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}
    
    fmt.Printf("len = %d | cap = %d\n\n", len(Brasil), cap(Brasil))
    for i := 0; i < 26; i++ {
        fmt.Println(Brasil[i])
    }
    
   
}
