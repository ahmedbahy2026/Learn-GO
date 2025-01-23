package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Birthdate string	`json:"birth_date"`
	CreatedAt time.Time `json:"created_at"`
}

type Admin struct {
	email string
	password string
	User
}

// (u user) --> called "Reciever" and used to attach a function (method) to a struct (user)
func (u User) OutputUserDetails(){
	// NOTE: u.firstName == (*u).firstName
	fmt.Println(u.FirstName, u.LastName, u.Birthdate)
}

func (u *User) ClearUserName(){
	u.FirstName = ""
	u.LastName = ""
}

func New(firstName, lastName, birthdate string) (*User,error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("firstName, lastName and birthdate are required")
	}

	return &User {
		FirstName: firstName,
		LastName: lastName,
		Birthdate: birthdate,
	},nil
}

func NewAdmin(email, password string) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password are required")
	}

	return &Admin {
		email: email,
		password: password,
		User: User{
			FirstName: "Ahmed",
			LastName: "Bahy",
			Birthdate: "08/16/2003",
			CreatedAt: time.Now(),
		},
	},nil
}

func (u User) Save() error {
	json, err := json.Marshal(u)

	if err != nil {
		return err
	}

	return os.WriteFile("file.json", json, 0644)
}