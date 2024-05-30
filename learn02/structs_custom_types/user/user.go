package user

import (
	"errors"
	"fmt"
	"time"
)
type User struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

func New(firstName string, lastName string, birthDate string) (appUser *User, err error){
	if (firstName == "" || lastName == "" || birthDate == ""){
		err = errors.New("empty values")
		return appUser, err
	}
	appUser = &User{
		firstName: firstName,
		lastName: lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}
	return appUser, nil
}

func (u User) OutputUseDetails(){
	fmt.Printf("Fist: %v\n", u.firstName)
	fmt.Printf("Last: %v\n", u.lastName)
	fmt.Printf("birthDate: %v\n", u.birthDate)
}

//setter method should has * , i.e. pass by reference
func (u *User) ClearUserName(){
	u.firstName = ""
	u.lastName = ""
}

type Admin struct {
	email string
	password string
	User
}

func NewAdmin(email, password string) Admin{
	return Admin{
		email: email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName: "test",
			birthDate: "03/11/1998",
			createdAt: time.Now(),
		},
	}

}