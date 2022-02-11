package main

import "fmt"

// Numa declaração de constantes, o identificador iota representa números sequenciais.

const (
	a = iota * 10
	b
	c
	d
	e
	f
	g
)

func main() {
	fmt.Println(a, b, c, e, f, g)
}
