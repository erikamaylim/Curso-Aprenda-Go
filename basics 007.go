package main

import "fmt"

func main() {
    // usando Sprint (string print)
	x := "oi"
	y := "tudo bem?"
	z := fmt.Sprint(x, ", ", y)
	fmt.Print(z)

}
