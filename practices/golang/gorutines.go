package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	// Método en una gorutine y se aplica el concepto de concurrencia.
	go nombre_lentoo("Jairo")

	var wait string
	fmt.Scanln(&wait)

}

func nombre_lentoo(name string) {
	letras := strings.Split(name, "")

	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
