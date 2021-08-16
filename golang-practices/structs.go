package main

import "fmt"

// Estructura
type User struct {
	edad             int
	nombre, apellido string
}

func main() {
	// Inicializando todas las propiedades de acuerdo al orden definido en la creació de la estructura.
	user := User{14, "Jairo", "Cuartas"}

	fmt.Println(user)
	fmt.Println(user.edad)

	// Inicializando propiedades especificas.
	user = User{nombre: "Jairo"}

	fmt.Println(user)
	fmt.Println(user.apellido)

	// Puntero de una estructura User pasada como parametro.
	pun_user := new(User)

	fmt.Println(pun_user)
	fmt.Println(user.edad)
}
