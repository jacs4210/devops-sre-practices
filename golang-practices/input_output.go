package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Output with print

	fmt.Println("Hi world")
	fmt.Printf("Mi edad es %d\n", 30)
	// fmt.Print("Bye world")
	fmt.Println("Bye world")

	// Scanf

	var nombre string

	fmt.Println("Cual es tu nombre?")
	fmt.Scanf("%s", &nombre)
	fmt.Printf("Hola %s", nombre)

	// Buffer
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Cual es tu nombre?")
	nombre, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print("Hola ", nombre)
	}
}
