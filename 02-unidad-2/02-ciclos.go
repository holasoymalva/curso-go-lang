package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	nombres := []string{"Malva", "Jorge", "Ainara", "Fer", "Marvin", "Nicolas"}

	// "Malva", "Jorge", "Ainara", "Fer", "Marvin", "Nicolas"
	// 0         1        2         3     4          5

	for i, nombre := range nombres {
		fmt.Printf("El nombre %s tiene el indice %d\n", nombre, i)
	}

	numeros := []int{1, 2, 3, 4, 5, 6, 7, 9}

	for i, numero := range numeros {
		fmt.Printf("El numero %s tiene el indice %d\n", numero, i)
	}

	// For simulando ser un while
	valor := 1
	for valor <= 10 {
		fmt.Println(valor)
		valor++
	}
}
