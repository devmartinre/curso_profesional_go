package main

import "fmt"

func main() {

	var curso = "Curso Profesional de Go!"
	var version = 10

	{
		fmt.Println(curso)

		{
			var version = 5

			fmt.Println(version)
			fmt.Println(curso)
		}

		fmt.Println(version)
	}

	fmt.Println(curso)
}
