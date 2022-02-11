/* SLICES (cont3)
    - Fatiamento de slices */

package main
import "fmt"

func main() {

    marvel := []string {"Homem-Aranha", "X-Men", "Guardiões da Galáxia", "Os Vingadores", "Shang-Chi"}
    fatia1 := marvel[0:2]
    fatia2 := marvel[2:]
    fatia3 := marvel[:3]
    fatia4 := marvel[:]
    fmt.Println(fatia1)
    fmt.Println(fatia2)
    fmt.Println(fatia3)
    fmt.Println(fatia4)
    for i := 0; i < len(marvel); i++ {
        fmt.Println(marvel[i])
    }
}