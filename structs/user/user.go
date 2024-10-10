package user

import (
	"errors"
	"fmt"
	"time"
)

// if this type needs to be used by another package,
// we need to capitalize the name "User" also field name "FirstName", we can keep it as "firstName" not to expose to other package on purpose
type User struct {
	FirstName   string
	LastName    string
	BirthDate   string
	CreatedAt   time.Time
	secretField string
}

type Admin struct {
	email    string
	password string
	User     // this is called anonymous embedding. `User: User` will be explicit embedding
}

// receiver argument: parenthesis before function name. It is used to turn the function into methods and attach it to the type that is specified in receiver
// func (userStruct *user) outputUserData() { <------- or this can work if we want to play with the original value
func (userStruct User) OutputUserData() {
	fmt.Println(userStruct.FirstName, userStruct.LastName, userStruct.BirthDate, userStruct.CreatedAt, userStruct.secretField)
}

// receiver argument makes a copy of the value of type by default
// so if we want to work with the original values, need to pass the pointer
func (userStruct *User) ClearUserName() {
	userStruct.FirstName = ""
	userStruct.LastName = ""
}

// [pattern] creation / construct function to create struct. new is a convention keyword for construct function
// [case 1] we can return with pointer

// [note] using only "New" is common for the constructor function that is stored in another package
// so instead of func NewUser:
func New(userFirstName, userLastName, userBirthDate string) (*User, error) {
	if userFirstName == "" || userLastName == "" || userBirthDate == "" {
		return nil, errors.New("first name, last name, birth date are required")

	}
	return &User{
		FirstName:   userFirstName,
		LastName:    userLastName,
		BirthDate:   userBirthDate,
		CreatedAt:   time.Now(),
		secretField: "this is secret field and can not be accessed from the other package",
	}, nil
}

// [case 2] or without and creates a copy
//
//	func NewUser(userFirstName, userLastName, userBirthDate string) user {
//		return user{
//			FirstName: userFirstName,
//			LastName:  userLastName,
//			BirthDate: userBirthDate,
//			CreatedAt: time.Now(),
//		}
//	}

func (admin *Admin) ShowAdminCredentials() {
	fmt.Println(admin.password)
	fmt.Println(admin.email)
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			FirstName: "NEW",
			LastName:  "ADMIN",
			BirthDate: "--",
			CreatedAt: time.Now(),
		},
	}
}
