package main

import "fmt"

func main() {
	monedas := [...]string{0: "Dólar Canadiense", 1: "Peso Méxicano", 2: "Dólar", 5: "Euro"}

	subMonedas := monedas[0:3] //slice

	fmt.Println(subMonedas)

	// fmt.Println(monedas[0])
	// fmt.Println(monedas[1])
	// fmt.Println(monedas[2])
	// fmt.Println(monedas[3])
	// fmt.Println(monedas[4])
	// fmt.Println(monedas[5])
}
