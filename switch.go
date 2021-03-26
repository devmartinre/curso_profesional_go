package main

import "fmt"

func main() {
	var calificacion int

	fmt.Print("Ingresa una califiación:")
	fmt.Scanf("%d", &calificacion)

	// switch {
	// case calificacion == 10:
	// 	fmt.Println("Felicidades, obtuviste una calificación perfecta.")
	// case calificacion == 8 || calificacion == 9:
	// 	fmt.Println("Aprobaste la materia.")
	// case calificacion == 6 || calificacion == 7:
	// 	fmt.Println("Aprobaste la materia, pero necesitas estudiar más.")
	// case calificacion >= 0 && calificacion <= 5:
	// 	fmt.Println("No aprobaste la materia.")
	// default:
	// 	fmt.Println("Calificación no válida.")
	// }

	switch calificacion {
	case 10:
		fmt.Println("Felicidades, obtuviste una calificación perfecta.")
	case 8, 9:
		fmt.Println("Aprobaste la materia.")
	case 6, 7:
		fmt.Println("Aprobaste la materia, pero necesitas estudiar más.")
	case 1, 2, 3, 4, 5:
		fmt.Println("No aprobaste la materia.")
	default:
		fmt.Println("Calificación no válida.")
	}
}
