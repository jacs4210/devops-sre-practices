package main

import "fmt"

func main() {
	slice := make([]int, 2, 2)
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 1)
	fmt.Println(cap(slice))
}
