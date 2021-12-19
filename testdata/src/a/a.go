package a

import "fmt"

type User struct {
	name string
}

func NewUser(name string) User {
	return User{
		name: name,
	}
}

func (u *User) SetName(name string) {
	u.name = name
}

func invalidFunc1() {
	u := User{}
	fmt.Println(u)
}
