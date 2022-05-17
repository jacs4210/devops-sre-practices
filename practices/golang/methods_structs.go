package main

import "fmt"

type User struct {
	edad             int
	nombre, apellido string
}

// Duplica el objeto User en tiempo de ejecución.
func (this User) nombre_completo() string {
	return this.nombre + " " + this.apellido
}

// Utiliza el objeto User que está asignado a la dirección en memoria *User, tiene mejor rendimiento en estructuras mas grandes y complejas.
func (this *User) set_name(n string) {
	this.nombre = n
}

func main() {
	user := new(User)

	user.set_name("Jairo")

	fmt.Println(user.nombre)

	user.apellido = "Cuartas"

	fmt.Println(user.nombre_completo())
}
