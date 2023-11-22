package models

import (
	"errors"
	"fmt"
)

type User struct {
	UserID    string `json:"userId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (u *User) LastFirstName() string {
	return fmt.Sprintf("%s, %s", u.LastName, u.FirstName)
}

func GetUserByID(userID string, users []User) (*User, error) {
	for _, user := range users {
		if user.UserID == userID {
			return &user, nil
		}
	}

	return nil, errors.New("user not found")
}
