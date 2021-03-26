package main

import "fmt"

func main() {
	usuarios := make(map[string]string)

	usuarios["eduardo"] = "eduardo@codigofacilito.com"

	if correo, ok := usuarios["eduardo"]; ok {
		fmt.Println(correo)
	} else {
		fmt.Println("No fue posible obtener el valor")
	}
}
