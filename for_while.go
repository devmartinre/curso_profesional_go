package main

import "fmt"

func main() {
	numero := 12875
	contador := 0

	for numero > 0 {

		numero = numero / 10
		contador++
	}

	fmt.Println("La cantidad de digitos es:", contador)
}
