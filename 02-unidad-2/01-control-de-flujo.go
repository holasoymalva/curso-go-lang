package main

import "fmt"

func main() {
	x := 10

	if x < 5 {
		fmt.Println("Valor menor a 5")
	} else {
		fmt.Println("Valor mayor a 5")
	}

	switch x {
	case 1:
		fmt.Println("X tiene valor 1")
	case 2:
		fmt.Println("X tiene valor 2")
	default:
		fmt.Println("X es mayor a 2")
	}
}

// if condition {
// 	// Codigo
// }

// if condition {
// 	// Codigo
// } else {
// 	// Codigo
// }

// switch condition{
// 	case valor1:
// 		// Codigo
// 	case valor2:
// 		// Codigo
// 	default:
// 		//codigo
// }
