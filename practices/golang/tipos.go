package main

import (
	"fmt"
	"strconv"
)

func main() {
	edad := 29
	edad_str := strconv.Itoa(edad)

	edad_ := "30"
	edad_int, err := strconv.Atoi(edad_)

	fmt.Println("Mi edad es " + edad_str)
	fmt.Println(edad_int)
	fmt.Println(err)
}
