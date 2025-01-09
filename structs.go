// package main

// import (
// 	"fmt"
// 	"time"
// )

// type user struct {
// 	firstName string
// 	lastName  string
// 	birthDate string
// 	createdAt time.Time
// }

// func main() {
// 	userFirstName := getUserData("Please enter your first name: ")
// 	userLastName := getUserData("Please enter your last name: ")
// 	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

// 	var appUser user

// 	appUser = user{
// 		firstName: userFirstName,
// 		lastName:  userLastName,
// 		birthDate: userBirthdate,
// 		createdAt: time.Now(),
// 	}


// 	outputUserDetails(appUser)
// }

// func outputUserDetails(val user) {
// 	// ...
// 	fmt.Println(val.firstName, val.lastName, val.birthDate)
// }

// func getUserData(promptText string) string {
// 	fmt.Print(promptText)
// 	var value string
// 	fmt.Scan(&value)
// 	return value
// }

//structs and pointers
// package main

// import (
// 	"fmt"
// 	"time"
// )

// type user struct {
// 	firstName string
// 	lastName  string
// 	birthDate string
// 	createdAt time.Time
// }

// func main() {
// 	userFirstName := getUserData("Please enter your first name: ")
// 	userLastName := getUserData("Please enter your last name: ")
// 	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

// 	var appUser user

// 	appUser = user{
// 		firstName: userFirstName,
// 		lastName:  userLastName,
// 		birthDate: userBirthdate,
// 		createdAt: time.Now(),
// 	}


// 	outputUserDetails(&appUser)
// }

// func outputUserDetails(val *user) {
// 	// ...
// 	fmt.Println((*val).firstName, val.lastName, val.birthDate)
// }

// func getUserData(promptText string) string {
// 	fmt.Print(promptText)
// 	var value string
// 	fmt.Scan(&value)
// 	return value
// }

//struct methods
// package main

// import (
// 	"fmt"
// 	"time"
// )

// type user struct {
// 	firstName string
// 	lastName  string
// 	birthDate string
// 	createdAt time.Time
// }
// func (val user) outputUserDetails() {
// 	// ...
// 	fmt.Println(val.firstName, val.lastName, val.birthDate)
// }

// func main() {
// 	userFirstName := getUserData("Please enter your first name: ")
// 	userLastName := getUserData("Please enter your last name: ")
// 	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

// 	var appUser user

// 	appUser = user{
// 		firstName: userFirstName,
// 		lastName:  userLastName,
// 		birthDate: userBirthdate,
// 		createdAt: time.Now(),
// 	}


// 	appUser.outputUserDetails()
// }



// func getUserData(promptText string) string {
// 	fmt.Print(promptText)
// 	var value string
// 	fmt.Scan(&value)
// 	return value
// }

//struct mutations

// package main

// import (
// 	"fmt"
// 	"time"
// )

// type user struct {
// 	firstName string
// 	lastName  string
// 	birthDate string
// 	createdAt time.Time
// }
// func (val *user) outputUserDetails() {
// 	// ...
// 	fmt.Println(val.firstName, val.lastName, val.birthDate)
// }

// func (val *user) clearUserName(){
// 	val.firstName=""
// 	val.lastName=""
// }
// func main() {
// 	userFirstName := getUserData("Please enter your first name: ")
// 	userLastName := getUserData("Please enter your last name: ")
// 	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

// 	var appUser user

// 	appUser = user{
// 		firstName: userFirstName,
// 		lastName:  userLastName,
// 		birthDate: userBirthdate,
// 		createdAt: time.Now(),
// 	}


// 	appUser.outputUserDetails()
// 	appUser.clearUserName()
// 	appUser.outputUserDetails()
// }



// func getUserData(promptText string) string {
// 	fmt.Print(promptText)
// 	var value string
// 	fmt.Scan(&value)
// 	return value
// }


//creation or constructor function

// package main

// import (
// 	"fmt"
// 	"time"
// )

// type user struct {
// 	firstName string
// 	lastName  string
// 	birthDate string
// 	createdAt time.Time
// }
// func (val *user) outputUserDetails() {
// 	// ...
// 	fmt.Println(val.firstName, val.lastName, val.birthDate)
// }

// func (val *user) clearUserName(){
// 	val.firstName=""
// 	val.lastName=""
// }

// func newUser(firstName,lastName,birthDate string) *user{
// 		return &user{
// 			firstName: firstName,
// 			lastName: lastName,
// 			birthDate: birthDate,
// 			createdAt: time.Now(),
// 		}
// }

// func main() {
// 	userFirstName := getUserData("Please enter your first name: ")
// 	userLastName := getUserData("Please enter your last name: ")
// 	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

// 	var appUser *user

// 	appUser = newUser(userFirstName,userLastName,userBirthdate)


// 	appUser.outputUserDetails()
// 	appUser.clearUserName()
// 	appUser.outputUserDetails()
// }



// func getUserData(promptText string) string {
// 	fmt.Print(promptText)
// 	var value string
// 	fmt.Scan(&value)
// 	return value
// }

// //validation

// package main

// import (
// 	"fmt"
// 	"time"
// 	"errors"
// )

// type user struct {
// 	firstName string
// 	lastName  string
// 	birthDate string
// 	createdAt time.Time
// }
// func (val *user) outputUserDetails() {
// 	// ...
// 	fmt.Println(val.firstName, val.lastName, val.birthDate)
// }

// func (val *user) clearUserName(){
// 	val.firstName=""
// 	val.lastName=""
// }

// func newUser(firstName,lastName,birthDate string) (*user,error){
// 	if firstName=="" || lastName=="" || birthDate==""{
// 		return nil,errors.New("All Details must be entered!")
// 	}
// 		return &user{
// 			firstName: firstName,
// 			lastName: lastName,
// 			birthDate: birthDate,
// 			createdAt: time.Now(),
// 		},nil 
// }

// func main() {
// 	userFirstName := getUserData("Please enter your first name: ")
// 	userLastName := getUserData("Please enter your last name: ")
// 	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

// 	var appUser *user

// 	appUser,err:= newUser(userFirstName,userLastName,userBirthdate)
// 	if err!=nil{
// 		fmt.Println(err)
// 		return 
// 	}

// 	appUser.outputUserDetails()
// 	appUser.clearUserName()
// 	appUser.outputUserDetails()
// }

// func getUserData(promptText string) string {
// 	fmt.Print(promptText)
// 	var value string
// 	fmt.Scanln(&value)
// 	return value
// }
// allias 
package main 

import "fmt"

type str string

func (text str) log(){
	fmt.Print(text)
}
func main(){
	var name str="alpha"
	name.log()
}

