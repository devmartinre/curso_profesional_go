package main

import "fmt"

type User struct {
	Name  string // Atributos
	Email string
}

func (u *User) setName(name string) {
	u.Name = name
}

func (u *User) getName() string {
	return u.Name
}

func main() {
	// var martin User

	// martin.Name = "Martin"
	// martin.Email = "martin@mail.com"
	// martin := User{Name: "martin", Email: "martin@mail.com"}

	// Name := "martin"
	// Email := "martin@mail.com"

	// martin := User{Name, Email}

	// fmt.Println(martin.Name)
	// fmt.Println(martin.Email)

	// fmt.Println(martin)

	martin := User{Name: "Martin", Email: "martin@mail.com"}

	martin.setName("Jose")

	fmt.Println(martin.getName())
}
