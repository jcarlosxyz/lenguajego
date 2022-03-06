package main

import "fmt"

func suma(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}
func variosRetornos(a int) (int, int, int) {

	return 2 * a, 3 * a, 4 * a
}
func variable_muchos(a int) (doble int, triple int, qua int) {
	doble = 2 * a
	triple = 3 * a
	qua = 4 * a
	return
}

//funcion con multiples retornos
func main() {
	fmt.Printf("La suma generada es %d\n", suma(1, 2))
	fmt.Printf("La suma generada es %d\n", suma(1, 2, 3))
	fmt.Printf("La suma generada es %d\n", suma(1, 2, 3, 4))
	fmt.Println(variosRetornos(2))
	fmt.Println("--------------multiple salidas--------")
	fmt.Println(variable_muchos(2))

}
