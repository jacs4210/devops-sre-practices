package main

import "fmt"

func main() {
	arreglo := [3]int{1, 2, 3}

	// Solo tomo los elementos despues de la primera posición.
	slice := arreglo[1:]

	// Solo tomo los elementos sin contar la ultima posición.
	slice = arreglo[:3]

	// Solo tomo los elementos que están dentro de las posiciones indicadas, en caso de que no existan valores entre ambas posiciones, se toma el ultimo valor de la posición final.
	slice = arreglo[1:2]

	fmt.Println(slice)
}
