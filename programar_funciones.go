package main

import (
	"fmt"
	"os"
)

// func funcion1() {
// 	fmt.Println("Hola, soy la primera función!!")
// }

// func funcion2() {
// 	fmt.Println("Hola, soy la segunda función!!")
// }

func main() {
	//defer funcion1()
	//funcion2()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Ups! al parecer el programa no finalizo correctamente.")
		}
	}()

	if file, err := os.Open("examples.txt"); err != nil {
		panic("No fue posible obtener el archivo!!")
	} else {
		defer func() {
			fmt.Println("Cerramos el archivo!!")
			file.Close()
		}()

		contenido := make([]byte, 254)

		longitud, _ := file.Read(contenido)

		contenidoArchivo := string(contenido[0:longitud])

		fmt.Println(contenidoArchivo)
	}
}
