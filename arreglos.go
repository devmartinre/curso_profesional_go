package main

import "fmt"

func main() {
	//var numeros [5]int

	//numeros := [5]int{100, 200, 300, 400, 500}
	numeros := [...]int{100, 200, 300, 400, 500}

	/*
		numeros[0] = 100
		numeros[1] = 200
		numeros[2] = 300
		numeros[3] = 400
		numeros[4] = 500
	*/

	fmt.Println(numeros)
}
