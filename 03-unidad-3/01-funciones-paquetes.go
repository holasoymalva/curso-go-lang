package main

import "fmt"

// Funciones Algebraicas
func sum(a int, b int) int {
	var resultado int
	resultado = a + b
	return resultado
}

func mult(a int, b int) int {
	var resultado int
	resultado = a * b
	return resultado
}

// Algoritmo de Numero Minimo

func minmax(a, b, c int) (int, int) {
	a = 23
	b = 33
	c = 57
	var min, max int

	if a > b && a > c {
		max = a
		if b < c {
			min = b
		} else {
			min = c
		}
	} else if b > a && b > c {
		max = b
		if a < c {
			min = a
		} else {
			min = c
		}
	} else {
		max = c
		if a < b {
			min = a
		} else {
			min = b
		}
	}
	return min, max
}

func main() {
	min, max := minmax(13, 4, 56)
	fmt.Printf("Min: %d \nMax: %d \n", min, max)
}
