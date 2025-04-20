package main

import (
	"fmt"
	"strconv"

	"github.com/Kgl-01/Structs_Go.git/user"
)

func main() {
	userFirstName := readUserInput("Please enter your first name: ")
	userLastName := readUserInput("Please enter your last name: ")
	userBirthDate := readUserInput("Please enter your birthdate (DD/MM/YYYY): ")
	userAge := readUserInput("Please enter your age: ")

	parsedUserAge, _ := strconv.Atoi(userAge)

	appUser, err := user.NewUser(userFirstName, userLastName, userBirthDate, parsedUserAge)

	if err == nil {
		// fmt.Print(appUser)
		admin := user.NewAdmin("karthik.gowda@consilio.com", "143.Amma")
		admin.OutputUserDetails()
		// appUser.OutputUserDetails()
		// appUser.ClearUserName()
		appUser.OutputUserDetails()
		return
	} else {
		fmt.Println(err)
	}
}

func readUserInput(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
