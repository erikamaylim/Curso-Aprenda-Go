/* MAPS */
package main
import "fmt"

func main() {
    amigos := map[string] int {
        "Amanda": 9999317,
        "Bruno": 9954242,
    }
    fmt.Println(amigos)
    fmt.Println(amigos["Amanda"])
    amigos["Hugo"] = 9888888
    fmt.Println(amigos)
    
    if será, ok := amigos["Rafael"]; !ok {  //comma ok idiom
        fmt.Println("não existe")
    } else {
        fmt.Println(será)
    }
}

/* Observações:
    - Utiliza o formato key:value.
    - E.g. nome e telefone
    - Performance excelente para lookups.
    - map[key]value{ key: value }
    - Acesso: m[key]
    - Key sem value retorna zero. Isso pode trazer problemas.
    - Para verificar: comma ok idiom.
        - v, ok := m[key]
        - ok é um boolean, true/false
    - Na prática: if v, ok := m[key]; ok { }
    - Para adicionar um item: m[v] = value
    - Maps *não tem ordem. */