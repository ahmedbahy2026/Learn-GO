package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.New(firstName, lastName, birthdate)
	// appAdmin, err := user.NewAdmin("ahmed@gmail.com", "********")
	if err != nil {
		fmt.Println(err)
		return
	}

	// ... do something awesome with that gathered data!

	// appAdmin.OutputUserDetails() == appAdmin.User.OutputUserDetails()
	// appAdmin.OutputUserDetails()
	// appAdmin.User.ClearUserName()
	// appAdmin.User.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
	appUser.Save()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	// Scan()   doesn't stop waiting for input when we enter newLine (press ENTER)
	// Scanln()         stop waiting for input when we enter newLine (press ENTER)
	// fmt.Scan(&value)
	fmt.Scanln(&value)
	return value
}
