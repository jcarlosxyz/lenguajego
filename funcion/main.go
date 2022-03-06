package main

import "fmt"

func main() {
	x := 10
	y := func() int {
		return x * 2
	}()
	fmt.Printf("El valor de la operacion es %d    ", y)
	fmt.Println(" ")
}
