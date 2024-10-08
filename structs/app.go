package main

import (
	"fmt"
	"time"
)

// if this type needs to be used by another package,
// we need to capitalize the name "User"
type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	var userFirstName string
	var userLastName string
	var userBirthDate string

	fmt.Println(time.Now())
	createUserData("enter your first name", &userFirstName)
	createUserData("enter your last name", &userLastName)
	createUserData("enter your birth date", &userBirthDate)

	var appUser user
	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthDate,
		createdAt: time.Now(),
	}

	// appUserShorthand := user{
	// 	userFirstName,
	// 	userLastName,
	// 	userBirthDate,
	// 	time.Now(),
	// }

	outputUserData(appUser)
}

func outputUserData(userStruct user) {
	fmt.Println(userStruct.firstName, userStruct.lastName, userStruct.birthDate, userStruct.createdAt)
}
func createUserData(promptText string, data *string) {
	fmt.Println(promptText)
	fmt.Scanln(data)

}
