package main

import "fmt"

type Persona struct {
	nombres string
	edad    int
}

type Empleado struct {
	id int
}
type FultEmpleado struct {
	Persona
	Empleado
}

func main() {

	fmt.Println("Progarma de herencia ")
	empleado1 := FultEmpleado{}
	empleado1.nombres = "nombre"
	empleado1.edad = 25
	empleado1.id = 1
	fmt.Println("%v", empleado1)
}
