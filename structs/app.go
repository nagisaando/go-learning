package main

import (
	"fmt"
	"time"

	"example.com/structs/user"
)

func main() {
	var userFirstName string
	var userLastName string
	var userBirthDate string

	fmt.Println(time.Now())
	createUserData("enter your first name", &userFirstName)
	createUserData("enter your last name", &userLastName)
	createUserData("enter your birth date", &userBirthDate)

	// [case 1] creates a construct manually

	// 1,
	// 	type user struct {
	// 	firstName string
	// 	lastName  string
	// 	birthDate string
	// 	createdAt time.Time
	// }

	//  var appUser user

	// appUser = user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthDate: userBirthDate,
	// 	createdAt: time.Now(),
	// }

	// 2,
	// appUserShorthand := user{
	// 	userFirstName,
	// 	userLastName,
	// 	userBirthDate,
	// 	time.Now(),
	// }

	// 3, with package

	// var appUser *user

	// 	appUser = &user.User{
	// 	FirstName: userFirstName,
	// 	LastName:  userLastName,
	// 	BirthDate: userBirthDate,
	// 	CreatedAt: time.Now(),
	// }

	// // [case 2] use construct function

	var appUser *user.User

	appUser, err := user.New(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputUserData()
	appUser.ClearUserName()
	appUser.OutputUserData()
}

// [example] use it as regular function:
//
//	func outputUserData(userStruct *user) {
//		// structs allows to access the value without dereferencing it. e.g. (*userStruct).firstName
//		fmt.Println((*userStruct).firstName)
//		fmt.Println(userStruct.firstName, userStruct.lastName, userStruct.birthDate, userStruct.createdAt)
//	}
//  outputUserData(&appUser)

func createUserData(promptText string, data *string) {
	fmt.Println(promptText)
	fmt.Scanln(data)

}
