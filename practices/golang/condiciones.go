package main

import "fmt"

/*
	== Igual a
	!= Diferente de
	<	Menor que
	>	Mayor que
	<=	Menor o igual que
	>=	Mayor o igual que
	&&	AND
	||	OR

*/
func main() {
	var num1 int
	var num2 int

	fmt.Println("====== Validar si el primer número es mayor que el segundo ====== ")
	fmt.Println("Ingrese el primer número a comparar")
	fmt.Scanln(&num1)
	fmt.Println("Ingrese el segundo número a comparar")
	fmt.Scanln(&num2)

	if num1 > num2 {
		fmt.Printf("%d es mayor que %d \n", num1, num2)
	} else if num1 < num2 {
		fmt.Printf("%d es menor que %d \n", num1, num2)
	} else {
		fmt.Printf("Son iguales \n", num1, num2)
	}
}
