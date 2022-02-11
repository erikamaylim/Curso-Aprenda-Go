package main
import "fmt"

/* A instrução continue é usada quando você quer pular uma parte restante do loop,
retornar ao topo desse loop e dar continuidade a uma nova iteração. */

func main() {
    for x := 0; x <= 20; x ++ {
        if x % 2 != 0 {
            continue
        }
        fmt.Println(x)
    }
}    