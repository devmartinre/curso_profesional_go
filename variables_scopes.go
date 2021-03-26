package main

import "fmt"

func modificarVariable(parametro string) {
	fmt.Println("Valor actual:", parametro)
	parametro = "Cambio de valor"
	fmt.Println("Nuevo valor:", parametro)

	fmt.Printf("%p\n", &parametro)
}

func main() {
	var curso = "Curso profesional de Go!"

	modificarVariable(curso)

	fmt.Println(curso)
	fmt.Printf("%p\n", &curso)
}
