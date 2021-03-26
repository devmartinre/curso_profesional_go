package main

import "fmt"

func promedio(califiaciones ...int) int {

	var promedio, suma int

	for _, calificacion := range califiaciones {
		suma = suma + calificacion
	}

	promedio = suma / len(califiaciones)

	return promedio
}

func main() {
	resultado := promedio(10, 10, 10, 9, 10)

	fmt.Println("Promedio:", resultado)
}
