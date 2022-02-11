/* Slice: slice multi-dimensional */
package main
import "fmt"

func main() {
    sliceofslices := [][]int {
        []int {1, 2, 3,},
        []int {4, 5, 6,},
        []int {7, 8, 9,},
    }
    fmt.Println(sliceofslices)
    
    outro := [][][][]int{
        [][][]int{
            [][]int {
                []int {1, 2, 3, 4},
                []int {5, 6, 7, 8},
            },
            [][]int {
                []int {9, 10, 11, 12},
                []int {13, 14, 15, 16},
            },
        },
    }
    fmt.Println(outro[0][1][0][2])
}