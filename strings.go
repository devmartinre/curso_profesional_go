package main

import (
	"fmt"
	"reflect"
)

func main() {

	//var curso string = "Curso profesional de Go!"
	//var curso = "Curso profesional de Go!"
	curso := "Curso profesional de Go!"

	longitud := len(curso)

	fmt.Println(longitud)

	primerCaracter := curso[0]
	ultimoCaracter := curso[len(curso)-1]

	fmt.Println(primerCaracter)
	fmt.Println(reflect.TypeOf(primerCaracter))

	fmt.Printf("%c\n", primerCaracter)
	fmt.Printf("%c\n", ultimoCaracter)
}
