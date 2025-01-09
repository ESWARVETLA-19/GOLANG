package user
import (
	"fmt"
	"time"
	"errors"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	CreatedAt time.Time
}
type Admin struct {
	Email string
	Password string
	Auser User
}
func (val *User) OutputUserDetails() {
	// ...
	fmt.Println((*val).FirstName, val.LastName, val.BirthDate)
}

func (val *User) ClearUserName(){
	val.FirstName=""
	val.LastName=""
}

func NewAdmin(Email ,Password string) *Admin{
	return &Admin{
		Email: Email,
		Password: Password,
		Auser: User{
			FirstName: "alpha",
			LastName: "beta",
			BirthDate: "123",
			CreatedAt: time.Now(),
		},
	}
}

func NewUser(firstName,lastName,birthDate string) (*User,error){
	if firstName=="" || lastName=="" || birthDate==""{
		return nil,errors.New("all Details must be entered")
	}
		return &User{
			FirstName: firstName,
			LastName: lastName,
			BirthDate: birthDate,
			CreatedAt: time.Now(),
		},nil 
}
