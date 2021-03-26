package main

import "fmt"

func main() {
	func() {
		fmt.Printf("Hola mundo, desde una función sin nombre\n")
	}()

	miFuncion := func(username string) string {
		message := fmt.Sprintf("Hola %s, te saludamos desde una función sin nombre",
			username)

		return message
	}

	mensajeFinal := miFuncion("Martin")
	fmt.Println(mensajeFinal)
}
