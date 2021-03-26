package main

import "fmt"

func main() {
	/*
		var nombre string
		var apellido string
		var pais string
	*/

	// var nombre, apellido, pais string
	// var nombre, apellido, pais = "José", "Rangel", "México"

	nombre, apellido, pais := "José", "Rangel", "México"
	edad, altura := 26, 1.83

	fmt.Println(nombre, apellido, pais, edad, altura)
}
