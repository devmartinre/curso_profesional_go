package main

import "fmt"

func main() {
	//Un mapa es una coleccion desordenada de elementos
	usuarios := map[int]string{}

	usuarios[1] = "Usuario 1"
	usuarios[2] = "Usuario 2"
	usuarios[3] = "Usuario 3"
	usuarios[4] = "Usuario 4"

	for id, username := range usuarios {
		fmt.Println(id, username)
	}
}
