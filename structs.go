package main

import (
	"fmt"
	"strconv"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	age       int
	createdAt time.Time
}

func newUser(firstName string, lastName string, birthDate string, age int) *user {

	appUser := &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		age:       age,
		createdAt: time.Now(),
	}

	return appUser
}

func (u user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}
func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func main() {
	userFirstName := readUserInput("Please enter your first name: ")
	userLastName := readUserInput("Please enter your last name: ")
	userBirthDate := readUserInput("Please enter your birthdate (DD/MM/YYYY): ")
	userAge := readUserInput("Please enter your age: ")

	parsedUserAge, _ := strconv.Atoi(userAge)

	var appUser user

	appUser = newUser(userFirstName, userLastName, userBirthDate, parsedUserAge)

	// fmt.Print(appUser)
	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func readUserInput(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
