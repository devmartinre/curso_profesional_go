package main

import "fmt"

var username string

func funcion1() {
	username = "Funcion1"
	fmt.Println("Función 1:", username)
}

func funcion2() {
	username = "Funcion2"
	fmt.Println("Función 2:", username)
}

func main() {
	username = "Jose"

	funcion1()
	funcion2()

	fmt.Println(username)
}
