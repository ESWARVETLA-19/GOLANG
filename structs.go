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
package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user

	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(),
	}


	outputUserDetails(&appUser)
}

func outputUserDetails(val *user) {
	// ...
	fmt.Println((*val).firstName, val.lastName, val.birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}

