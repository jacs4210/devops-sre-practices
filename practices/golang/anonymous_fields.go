package main

import "fmt"

type Human struct {
	name string
}

func (this Human) hablar() string {
	return "Bla, bla, bla"
}

type Dummy struct {
	name string
}

type Tutor struct {
	Human
	Dummy
}

func (this Tutor) hablar() string {
	return "Hello"
}

func main() {
	tutor := Tutor{Human{"Jairo"}, Dummy{"Angie"}}

	fmt.Println(tutor.Dummy.name)
	fmt.Println(tutor.hablar())
}
