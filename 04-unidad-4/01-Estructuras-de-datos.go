package main

import "fmt"

type item struct {
	value interface{}
	next  *item
}

type Stack struct {
	top  *item
	size int
}

func (stack *Stack) Len() int {
	return stack.size
}

func (stack *Stack) Push(value interface{}) {
	stack.top = &item{
		value: value,
		next:  stack.top,
	}
	stack.size++
}

func (stack *Stack) Pop() (value interface{}) {
	if stack.Len() > 0 {
		value = stack.top.value
		stack.top = stack.top.next
		stack.size--
		return
	}
	return nil
}

func main() {
	// Arreglo
	// var miArreglo [10]int

	// fmt.Printf("%+v\n", miArreglo)

	// for i := 0; i < 10; i++ {
	// 	miArreglo[i] = i + 1
	// }

	// fmt.Printf("%+v\n", miArreglo)
	// fmt.Printf("La longitud del arreglo es %d", len(miArreglo))

	// Slice
	// miSlice := []int{}

	// for i := 0; i < 100; i++ {
	// 	miSlice = append(miSlice, i+1)
	// }

	// fmt.Printf("%+v", miSlice)

	//mapas

	// personas := make(map[string]int)

	// personas["Isai"] = 25
	// personas["Raquel"] = 21
	// personas["Ramiro"] = 106
	// personas["Gerardo"] = 19

	// for llave, valor := range personas {
	// 	fmt.Println(llave, " = ", valor)
	// }

	// delete(personas, "Ramiro")

	// for llave, valor := range personas {
	// 	fmt.Println(llave, " = ", valor)
	// }

	stack := new(Stack)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println(stack.Pop())
	fmt.Println(stack)
	fmt.Println(stack.Pop())
	fmt.Println(stack)
}
