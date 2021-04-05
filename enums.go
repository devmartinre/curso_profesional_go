package main

import "fmt"

type UserType int

const (
	Teacher UserType = 1
	Student UserType = 2
)

type User struct {
	Username string
	Type     UserType
}

func main() {
	jose := User{Username: "jose", Type: Student}
	martin := User{Username: "martin", Type: Teacher}

	fmt.Println(jose)
	fmt.Println(martin)

	if jose.Type == Student {
		fmt.Println("El usuario", jose.Username, "es un estudiante!")
	}
}
