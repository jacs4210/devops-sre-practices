package main

import "fmt"

func main() {

	// Array
	var array [2]int

	fmt.Println(array)

	// Matriz
	var matriz [2][2]int

	fmt.Println(matriz)

	// Longitud
	fmt.Println(len(array))

	// Inicialización de un arreglo
	array = [2]int{1, 2}

	fmt.Println(array)
}
