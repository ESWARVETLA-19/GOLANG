package main

import(
	"fmt"
)
//functions as values and types

type  transForm func(int)int //useful when its complex func type
func main(){
	numbers:=[]int{1,2,3,4}
	// a:=5
	fmt.Println(Transformnum(&numbers,db))
	fmt.Println(Transformnum(&numbers,trp))
	// fmt.Print(Transformer(a))
	
}
func Transformnum(numbers *[]int,transform transForm) []int{
	dnumbers:=[]int{}
	for idx,val:=range *numbers{
		dnumbers=append(dnumbers, transform(val*idx))
	}
	return dnumbers
}

func db(val int)int{
	return val*2
}

func trp(val int)int{
	return val*3
}

//returning functions as values
// func getFunc() func(int)int{
// 	return funcname  //we can use conditions 
// }

//Ananoymous function -use if its one off task 
// fc:=transForm(&numbers,func (number int)int{
// 	return number*2
// })

//closures-closures are functions that reference variables from their surrounding scope
// func Transformer(fact int) func(int) int{
// 	return func(num int)int{
// 		return num*fact
// 	}
// }

//varidic functions
// sum:=sumup(1,2,3,4,5,6,7,8,9)
// func sumup(numbers ...int) int{
// 	sum:=0
// 	sum+val
// 	return sum 
// }
//sumup(a int,numbers ...int)int{//take 1st element as a and remaing as list ...}

