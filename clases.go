package main

import "fmt"

//creamos la clase  empleados
//hola primer cambio
type Empledo struct {
	id     string
	nombre string
}

//funcion para inicializar el objeto 4 método
func NuevoEmpleado(id string, nombre string) *Empledo {

	return &Empledo{
		id:     id,
		nombre: nombre,
	}
}

func main() {

	fmt.Println("---Programa de clases--------")
	e := Empledo{} //declaramos el objeto
	e.id = "1"
	e.nombre = "juan carlos zarate molina"
	fmt.Println("---mostramos la clase--------")
	fmt.Println("%v", e)
	//segundo metoto decrear objeto
	e2 := Empledo{
		id:     "123456",
		nombre: "Entrada de datos",
	}
	fmt.Println("---mostramos la clase metodo 2--------")
	fmt.Print("%v\n", e2)
	fmt.Println("\n---mostramos la clase metodo 4--------")
	e4 := NuevoEmpleado("12", "juan")
	fmt.Println("%v\n", e4)
}
