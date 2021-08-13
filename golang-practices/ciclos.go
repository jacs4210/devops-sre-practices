package main

import "fmt"

func main() {

	// For con todos sus elementos
	for i := 0; i < 5; i++ {
		fmt.Println("Hi World complete")
	}

	// For solo con condición
	i := 0
	for i < 5 {
		fmt.Println("Hi World with condition")
		i++
	}

	// For sin elementos
	i = 0
	for {
		fmt.Println("Hi World without elements")
		i++

		if i > 5 {
			break
		}
	}

}
