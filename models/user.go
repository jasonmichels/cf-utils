package models

import "fmt"

type UserSession struct {
	UserID     string `json:"userID"`
	UserName   string `json:"userName"`
	UserEmail  string `json:"userEmail"`
	IsLoggedIn bool   `json:"isLoggedIn"`
}

type User struct {
	UserID    string `json:"userId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (u *User) LastFirstName() string {
	return fmt.Sprintf("%s, %s", u.LastName, u.FirstName)
}
