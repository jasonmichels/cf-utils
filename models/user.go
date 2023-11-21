package models

import "fmt"

type User struct {
	UserID    string `json:"userId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (u *User) LastFirstName() string {
	return fmt.Sprintf("%s, %s", u.LastName, u.FirstName)
}
