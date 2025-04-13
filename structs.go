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

func main() {
	userFirstName := readUserInput("Please enter your first name: ")
	userLastName := readUserInput("Please enter your last name: ")
	userBirthDate := readUserInput("Please enter your birthdate (DD/MM/YYYY): ")
	userAge := readUserInput("Please enter your age: ")

	parsedUserAge, _ := strconv.Atoi(userAge)

	var appUser user

	appUser = user{firstName: userFirstName, lastName: userLastName, birthDate: userBirthDate, age: parsedUserAge, createdAt: time.Now()}

	fmt.Print(appUser)
}

func outputUserDetails(u user) {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}

func readUserInput(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
