package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	copia := make([]int, 1)

	// La cantidad minima de elementos es igual a la longitud definida en el slice destino, en este caso, 1 solo elemento se copiará del slice fuente.
	copy(copia, slice)

	fmt.Println(slice)
	fmt.Println(copia)
}
