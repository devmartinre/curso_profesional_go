package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Active bool
}

type Student struct {
	User User
	Id   string
}

func main() {
	jose := User{Name: "Jose", Email: "martin@mail.com", Active: true}
	martin := User{Name: "Martin", Email: "jose@mail.com", Active: true}
	joseStudent := Student{User: jose, Id: "1f1"}

	fmt.Println(martin)
	fmt.Println(joseStudent.User.Name)
}
