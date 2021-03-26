package main

import "fmt"

func suma(numero1, numero2 int) (int, string) {
	return numero1 + numero2, "un mensaje desde la función suma"
}

func main() {
	resultado, mensaje := suma(10, 20)

	fmt.Println(resultado, mensaje)
}
