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
	fmt.Println(a[0][0], "\t\t", a[0][1], "\t\t", a[0][2], "\t\t")
	fmt.Println(a[1][0], "\t\t", a[1][1], "\t\t\t", a[1][2], "\t\t")
	fmt.Println(a[2][0], "\t", a[2][1], "\t\t", a[2][2], "\t\t")
	fmt.Println(a[3][0], "\t\t", a[3][1], "\t\t", a[3][2], "\t\t")

}
