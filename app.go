
//hello world
// package main
// import "fmt"
// func main(){
// 	fmt.Print("Hello World!")
// }

// Sprintf() -store formatted string 
// go mod init path (go build-./modulename executable file main)
/* building multi line string - `abc
                                  defghijklmno`   */


// calc program variables values
// package main
// import (
// 	"fmt"
//     "math"
// )
// func main(){
// 	var invest=1000;
// 	var expected=5.5;
// 	var years=10;

// 	var futureval=float64(invest) * math.Pow(1 + expected / 100,float64(years))
// 	fmt.Print(futureval)
// }

//functions

package main
import( "fmt")
func fact(a int) int{
	if(a==0){
		return 1;
	}else{
		return a*fact(a-1);
	}
}
func maain(){
	var n int;
	fmt.Scan(&n)
	fmt.Print(fact(n))
}

//multiple return func add(a,b int)(a+b int,a int)(implicit return) there is no need to return values just return is enough and no need to declare variables
// if else if else
// return stop execution
// for i:=0;i<n;i++{} 
// for{} infinite loop
/*
switch choice{
case 1:

case 2:

default:
	break// //return 
}*/


/// writing and reading in to files
// package main 
// import(
// 	"fmt"
//   "os"
// )
// func writeData(data int){
// 	filedata:=fmt.Sprint(data)
// 	os.WriteFile("file1.txt",[]byte(filedata),0644)
// }
// func main(){
// 	writeData(1000)
// }
//strconv 
// package main 
// import(
// 	"fmt"
//     "os"

// )
// const file="file1.txt";
// func writeData(data string){
	
// 	os.WriteFile(file,[]byte(data),0644)
// }
// func readData() string{
// 	dt,_:=os.ReadFile(file)
// 	return string(dt)

// }
// func main(){
// 	writeData("Alpha")
// 	fmt.Print(readData())
// }

// errors return empty data and second value as error 
// nil
//error.New("content") import "errors"
//panic("content") dont want continue

//split code across files in same packages
// package main

// import (
// 	"errors"
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// const accountBalanceFile = "balance.txt"

// func getBalanceFromFile() (float64, error) {
// 	data, err := os.ReadFile(accountBalanceFile)

// 	if err != nil {
// 		return 1000, errors.New("Failed to find balance file.")
// 	}

// 	balanceText := string(data)
// 	balance, err := strconv.ParseFloat(balanceText, 64)

// 	if err != nil {
// 		return 1000, errors.New("Failed to parse stored balance value.")
// 	}

// 	return balance, nil
// }

// func writeBalanceToFile(balance float64) {
// 	balanceText := fmt.Sprint(balance)
// 	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
// }

// func main() {
// 	var accountBalance, err = getBalanceFromFile()

// 	if err != nil {
// 		fmt.Println("ERROR")
// 		fmt.Println(err)
// 		fmt.Println("---------")
// 		// panic("Can't continue, sorry.")
// 	}

// 	fmt.Println("Welcome to Go Bank!")

// 	for {
// 		presentOptions()

// 		var choice int
// 		fmt.Print("Your choice: ")
// 		fmt.Scan(&choice)

// 		// wantsCheckBalance := choice == 1

// 		switch choice {
// 		case 1:
// 			fmt.Println("Your balance is", accountBalance)
// 		case 2:
// 			fmt.Print("Your deposit: ")
// 			var depositAmount float64
// 			fmt.Scan(&depositAmount)

// 			if depositAmount <= 0 {
// 				fmt.Println("Invalid amount. Must be greater than 0.")
// 				// return
// 				continue
// 			}

// 			accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
// 			fmt.Println("Balance updated! New amount:", accountBalance)
// 			writeBalanceToFile(accountBalance)
// 		case 3:
// 			fmt.Print("Withdrawal amount: ")
// 			var withdrawalAmount float64
// 			fmt.Scan(&withdrawalAmount)

// 			if withdrawalAmount <= 0 {
// 				fmt.Println("Invalid amount. Must be greater than 0.")
// 				continue
// 			}

// 			if withdrawalAmount > accountBalance {
// 				fmt.Println("Invalid amount. You can't withdraw more than you have.")
// 				continue
// 			}

// 			accountBalance -= withdrawalAmount // accountBalance = accountBalance + depositAmount
// 			fmt.Println("Balance updated! New amount:", accountBalance)
// 			writeBalanceToFile(accountBalance)
// 		default:
// 			fmt.Println("Goodbye!")
// 			fmt.Println("Thanks for choosing our bank")
// 			return
// 			// break
// 		}
// 	}
// }

// package main

// import (
// 	"fmt"

// 	"example/learning/fileops"
// 	"github.com/Pallinder/go-randomdata"
// )

// const accountBalanceFile = "balance.txt"

// func main() {
// 	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

// 	if err != nil {
// 		fmt.Println("ERROR")
// 		fmt.Println(err)
// 		fmt.Println("---------")
// 		// panic("Can't continue, sorry.")
// 	}

// 	fmt.Println("Welcome to Go Bank!")
// 	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

// 	for {
// 		presentOptions()

// 		var choice int
// 		fmt.Print("Your choice: ")
// 		fmt.Scan(&choice)

// 		// wantsCheckBalance := choice == 1

// 		switch choice {
// 		case 1:
// 			fmt.Println("Your balance is", accountBalance)
// 		case 2:
// 			fmt.Print("Your deposit: ")
// 			var depositAmount float64
// 			fmt.Scan(&depositAmount)

// 			if depositAmount <= 0 {
// 				fmt.Println("Invalid amount. Must be greater than 0.")
// 				// return
// 				continue
// 			}

// 			accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
// 			fmt.Println("Balance updated! New amount:", accountBalance)
// 			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
// 		case 3:
// 			fmt.Print("Withdrawal amount: ")
// 			var withdrawalAmount float64
// 			fmt.Scan(&withdrawalAmount)

// 			if withdrawalAmount <= 0 {
// 				fmt.Println("Invalid amount. Must be greater than 0.")
// 				continue
// 			}

// 			if withdrawalAmount > accountBalance {
// 				fmt.Println("Invalid amount. You can't withdraw more than you have.")
// 				continue
// 			}

// 			accountBalance -= withdrawalAmount // accountBalance = accountBalance + depositAmount
// 			fmt.Println("Balance updated! New amount:", accountBalance)
// 			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
// 		default:
// 			fmt.Println("Goodbye!")
// 			fmt.Println("Thanks for choosing our bank")
// 			return
// 			// break
// 		}
// 	}
// }

//pointers
// package main 
// import "fmt"
// func add(a *int,b int){
//     *a= *a+b
// }

// func maain(){
//     x:=5
//     y:=10
//     add(&x,y)
//     fmt.Println(x,y)
// }
// without pointers
// package main 
// import "fmt"
// func add(a,b int){
//     a= a+b
// }

// func main(){
//     x:=5
//     y:=10
//     add(x,y)
//     fmt.Println(x,y)
// }

// package main 
// import "fmt"
// func add(a int,b int){
//     // fmt.Println(&a)
//     a= a+b
// }

// func main(){
//     var x int=5
//     a:=&x
//     y:=10
//     add(x,y)
//     fmt.Println(a)
//     fmt.Println(*a,y)
// }