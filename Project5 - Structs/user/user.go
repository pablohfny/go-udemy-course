package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name and birthdate are required")
	}
	return &User{
		firstName,
		lastName,
		birthdate,
		time.Now(),
	}, nil
}

func (user User) OutputUserDetails() {
	fmt.Println("FirstName", user.firstName)
	fmt.Println("LirstName", user.lastName)
	fmt.Println("Birthdate", user.birthdate)
	fmt.Println("CreatedAt", user.createdAt)
}

func (user *User) ClearUserName() {
	user.firstName = ""
	user.lastName = ""
}

type Admin struct {
	email    string
	password string
	User
	// User User
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthdate: "---",
			createdAt: time.Now(),
		},
	}
}
