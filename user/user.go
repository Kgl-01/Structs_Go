package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	age       int
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User:     User{firstName: "Admin", lastName: "Admin", birthDate: "---", age: 26, createdAt: time.Now()},
	}
}

func NewUser(firstName string, lastName string, birthDate string, age int) (*User, error) {

	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("firstName, lastName and birthDate is missing")
	}

	appUser := &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		age:       age,
		createdAt: time.Now(),
	}

	return appUser, nil
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}
