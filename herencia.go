//majeno de las interfaces
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

func (fultEmpleado FultEmpleado) mensajePersona() string {
	return "Este es un fultEmpleado"
}

type Informacio interface {
	mensajePersona() string
}

func mensajePersona(p Informacio) {
	fmt.Println(p.mensajePersona())
}

func main() {
	//cambio de 2 de marzo
	fmt.Println("Progarma de herencia ")
	empleado1 := FultEmpleado{}
	empleado1.nombres = "nombre"
	empleado1.edad = 25
	empleado1.id = 1
	mensajePersona(empleado1)
	//fmt.Println("%v", empleado1)
}
