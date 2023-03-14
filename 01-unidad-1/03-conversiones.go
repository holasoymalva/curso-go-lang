package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 42
	var b string = string(a)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(a)
	fmt.Println(b)

	var c int = 42
	var d float32 = float32(c)
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(d))
	// Cual es el tipo de dato en golang
}
