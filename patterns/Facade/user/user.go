package user

import "fmt"

type User struct {
	Name     string
	LastName string
}

func (u *User) NewUser(name, lastName string) *User {
	fmt.Printf("user with name: %s %s has been created \n", name, lastName)
	return &User{
		Name:     name,
		LastName: lastName,
	}
}

// more code...
