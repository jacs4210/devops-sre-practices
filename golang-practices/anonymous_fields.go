package main

import "fmt"

type Human struct {
	name string
}

type Dummy struct {
	name string
}

type Tutor struct {
	Human
	Dummy
}

func main() {
	tutor := Tutor{Human{"Jairo"}, Dummy{"Angie"}}

	fmt.Println(tutor.Dummy.name)
}
