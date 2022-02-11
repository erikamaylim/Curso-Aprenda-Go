package main

import (
	"fmt"

)

var x = 10

func main() {

    // Aqui se usou o Printf, e não o Println
	fmt.Printf("%v, %T", x, x)
}
