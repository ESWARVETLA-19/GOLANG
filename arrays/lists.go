package main
import (
	"fmt"
)
// func main(){
// 	// arr:=[4]int64{1,2,3,4}
// 	// var arr [5]int 
// 	// arr=[5]int{1,2}
// 	var arr [5]int=[5]int{1,2}
// 	arr[3]=4
// 	ls:=arr[:1]
// 	fmt.Print(arr) 
// 	fmt.Print(len(ls),cap(ls))//1 5
// 	fmt.Print("ls:",ls)
// 	l:=ls[:5]
// 	fmt.Print("l:",l)
// 	//len exact size of the present array
// 	//cap gives the highest capacity that can store
// }


//Dynamic Array
func main(){
	var arr[] int 
	for i:=0;i<10;i++{
		arr=append(arr,i)
	}
	fmt.Print(arr)

}
