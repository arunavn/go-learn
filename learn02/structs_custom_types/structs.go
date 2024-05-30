package main

import (
	"fmt"

	"example.com/structs/user"
)


func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *(user.User)

	// appUser = User{
	// 	firstName: userFirstName,
	// 	lastName: userLastName,
	// 	birthDate: userBirthDate,
	// 	createdAt: time.Now(),
	// }

	// same
	// appUser = User{
	// 	userFirstName,
	// 	userLastName,
	// 	userBirthDate,
	// 	time.Now(),
	// }

	// appUser = User{} //Empty struct

	appUser, err := user.New(userFirstName, userLastName, userBirthDate)
	if err != nil{
		panic(err)
	}

	// ... do something awesome with that gathered data!
	appUser.OutputUseDetails()
	// appUser.ClearUserName()
	// appUser.OutputUseDetails()

	// fmt.Println(userFirstName, userLastName, userBirthDate)

	admin1 := user.NewAdmin("s", "d" )
	admin1.OutputUseDetails()
}


func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
