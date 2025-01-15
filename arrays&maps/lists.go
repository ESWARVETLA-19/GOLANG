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
// type Product struct{
// 	id int64 
// 	title string
// 	price float64
// }
// func main(){
// 	var arr[] int 
// 	for i:=0;i<10;i++{
// 		arr=append(arr,i)
// 	}
// 	fmt.Print(arr)	
// 	ls:=arr[:3]
// 	fmt.Print(ls)
// 	l:=ls[3:10]
// 	fmt.Print(l)
// 	darr:=[]string{"ai","ml"}
// 	darr[1]="dl"
// 	darr=append(darr,"golang")
// 	fmt.Print(darr)
// 	products:=[]Product{{01,"earphones",1000},{02,"sneakers",2000}}
// 	fmt.Print(products)
// 	ar:=[]string{"cp"}
// 	darr=append(darr,ar...)
// 	fmt.Print(darr)

// }

//maps
func main(){

	web:=map[string]string{
		"google":"google.com",
		"amazon":"aws.com",
		"lendi":"lendi.edu.in",	
	}
	web["linkedin"]="linkedin.com"
	fmt.Print(web)
	fmt.Print(web["google"])
	delete(web,"lendi")
	fmt.Print(web)
	for index,value:=range web{
		fmt.Print(index,value)
	}

}

//type allias 
// type floatMap map[string]float64

// func (m floatMap) output(){
// 	fmt.Print(m)
// }


// //make
// func main(){
// 	user:=make([]string,2,100)//type,length,capacity
// 	user[0]="alpha"
// 	fmt.Print(user)
// 	course:=make(floatMap,3)
// 	course["go"]=500.0
// 	course.output()
// 	//make(map[type]type,capacity)
// }
