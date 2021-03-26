package main

import "fmt"

func main() {

	meses := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio",
		"Agosto", "Septiembre"}

	// Un puntero
	// Una longitud
	// Una capacidad

	longitud := len(meses)
	capacidad := cap(meses)

	fmt.Printf("La longitud es: %v, la capacidad es: %v %p\n", longitud, capacidad, meses)

	meses = append(meses, "Octubre") // Si la estructura se encuentra es su capacidad maxima se genera un nuevo arreglo
	meses = append(meses, "Noviembre")
	meses = append(meses, "Diciembre")

	longitud = len(meses)
	capacidad = cap(meses)

	fmt.Printf("La longitud es: %v, la capacidad es: %v\n %p", longitud, capacidad, meses)
}
