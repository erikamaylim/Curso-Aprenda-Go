/*  - Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
    - "Nome", "Sobrenome", "Hobby favorito"
- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados. */
package main

import "fmt"

func main() {

	a := [][]string{
		[]string{"Nome:", "Sobrenome:", "Hobby:"},
		[]string{"John", "Snow", "Caçar Caminhantes Brancos"},
		[]string{"Daenerys", "Targaryen", "Criar dragões"},
		[]string{"Tyrion", "Lannister", "Beber vinho"},
	}
	
	 for _, v := range a {
        fmt.Println(v[0], "             ", v[1], "             ", v[2])

	}
}
