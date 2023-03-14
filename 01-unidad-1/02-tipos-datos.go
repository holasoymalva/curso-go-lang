package main

import "fmt"

func main() {
	// Booleanos
	var a bool = true
	var b bool = false
	fmt.Println(a)
	fmt.Println(b)

	// Numericos
	var c int8 = 125
	var d int16 = 32767
	var e int32 = 2145325676
	var f int64 = 53453453453453454
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	// Numericos sin signo
	var g uint8 = 125
	var h uint16 = 32767
	var i uint32 = 2145325676
	var j uint64 = 53453453453453454
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)

	// PuntoFlotante
	var k float32 = 3.14159
	var l float64 = 3.141592634503405345345
	fmt.Println(k)
	fmt.Println(l)

	// Cadenas
	var m string = "Ella no te ama"
	fmt.Println(m)

	// Byte y Rune
	var n byte = 'A'
	var o rune = '♥'
	fmt.Println(n)
	fmt.Println(o)
}
