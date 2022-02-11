package main

import (
	"fmt"
	"reflect"
)

var x = 10

func main() {

	fmt.Println(reflect.TypeOf(x), x)
}
