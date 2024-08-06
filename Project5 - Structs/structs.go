package main

import (
	"fmt"
	"structs/user"
)

type str string

// CAn add method to alias of string
func (text str) log() {
	fmt.Println(text)
}

func main() {
	dataFirstName := getUserData("Please enter your first name: ")
	dataLastName := getUserData("Please enter your last name: ")
	dataBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// user := User{
	// 	FirstName: dataFirstName,
	// 	LastName:  dataLastName,
	// 	Birthdate: dataBirthdate,
	// }

	newUser, err := user.New(dataFirstName, dataLastName, dataBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	newUser.OutputUserDetails()
	newUser.ClearUserName()
	newUser.OutputUserDetails()

	admin := user.NewAdmin("", "")

	admin.OutputUserDetails()
	// outputUserDetails(user)

	var name str
}

// func outputUserDetails(user User) {
// 	fmt.Println("FirstName", user.firstName)
// 	fmt.Println("LirstName", user.lastName)
// 	fmt.Println("Birthdate", user.birthdate)
// 	fmt.Println("CreatedAt", user.createdAt)
// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
