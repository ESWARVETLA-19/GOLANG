
package main

import (
	"fmt"
	"example/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User

	appUser,err:= user.NewUser(userFirstName,userLastName,userBirthdate)
	if err!=nil{
		fmt.Println(err)
		return 
	}
	admin:=user.NewAdmin("a1","pass")
	admin.Auser.OutputUserDetails()
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}



func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}