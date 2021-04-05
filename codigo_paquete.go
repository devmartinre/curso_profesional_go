package main

import (
	"fmt"

	"./CodigoFacilito"
)

func main() {
	curso := CodigoFacilito.Curso{Titulo: "Curso Profesional de Go!"}

	fmt.Println(curso.ObtenerTitulo())
}
