/* Slice: a surpresa do array subjacente */
package main
import "fmt"

func main() {
    slice1 := []int {1, 2, 3, 4, 5}
    fmt.Println("slice1:\t\t\t", slice1)
    slice2 := append(slice1[:2], slice1[3:]...)
    fmt.Println("slice2:\t\t\t", slice2)
    fmt.Println("slice1 após append:\t", slice1)
   
}

/* ATENÇÃO!!!

slice1:			     [1 2 3 4 5]
slice2:			     [1 2 4 5]
slice1 após append:	 [1 2 4 5 5]

slice2 utiliza o mesmo array subjacente que slice1. Ao executar o append, o array foi modificado,
alterando o resultado de slice1, o que não era a intenção.
 
Como evitar:
• Usar a mesma variável
• Copiar item por item usando for loop */
