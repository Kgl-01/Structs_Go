package main

import "fmt"

func main() {
	firstName := readUserInput("Please enter your first name: ")
	lastName := readUserInput("Please enter your last name: ")
	birthDate := readUserInput("Please enter your birthdate (DD/MM/YYYY): ")

	fmt.Println(firstName, lastName, birthDate)
}

func readUserInput(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
