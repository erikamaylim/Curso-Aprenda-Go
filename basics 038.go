/* SLICES (cont4)
    - Excluindo um elemento de uma slice com o uso de fatiamento */

package main
import "fmt"

func main() {

    marvel := []string {"Homem-Aranha", "X-Men", "Guardiões da Galáxia", "Os Vingadores", "Shang-Chi"}
    marvel = append(marvel[:2], marvel[3:]...)
    fmt.Println(marvel)
}