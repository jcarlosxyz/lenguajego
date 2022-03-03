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
	sueldo int
}
type MedioTiempoEmpleado struct {
	Persona
	Empleado
	MesesContrato int
}

///////////////////////////////////////
func (fultEmpleado FultEmpleado) mensajePersona() string {
	return "Este es un fultEmpleado"
}

func (medioTiempoEmpleado MedioTiempoEmpleado) mensajePersona() string {
	return "Este es un medio empleado"
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
	empleado1.sueldo = 25000
	mensajePersona(empleado1) //utilizacion de la interface
	//***********************************************
	empleado2 := MedioTiempoEmpleado{}
	empleado2.nombres = "nombre"
	empleado2.edad = 50
	empleado2.id = 2
	empleado2.MesesContrato = 10
	mensajePersona(empleado2)
	//fmt.Println("%v", empleado1)
}
