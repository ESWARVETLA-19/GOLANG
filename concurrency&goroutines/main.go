package main

import (
	"fmt"
	"time"
)

//concurrency- Running tasks in parallel
//in GoLang we use Goroutines(go keyword)
/*Goroutines are lightweight, concurrent execution threads 
in the Go programming language*/
//dispatch and exit -solution channels
//channels - value used to communication well create by using make(chan )
//go routins doesnt support retuning values

//error channels are used to catch the errors in go routines 
//(name chan error)
func greet(phrase string,done chan bool){
	fmt.Println("hello!",phrase)
	done<-true
}

func slowgreet(phrase string,done chan bool){
	time.Sleep(2*time.Second)
	fmt.Println("hello!",phrase)
	done <- true
}

// func main(){
// 	greet("alpha")
// 	go greet("beta")
// 	go slowgreet("charlie")
// 	greet("delta")
// 	slowgreet("echo")
// 	slowgreet("foxtrot")
// }


// func main(){
	
// 	done:=make(chan bool)
// 	go slowgreet("charlie",done)
// 	<-done//after this signal get will move forward
// 	greet("alpha")
// }

func main(){
	done:=make([]chan bool,6)
	for i,_:=range done{
		done[i]=make(chan bool)
	}
	go greet("alpha",done[0])
	go greet("beta",done[1])
	go slowgreet("charlie",done[2])
	go greet("delta",done[3])
	go slowgreet("echo",done[4])
	go slowgreet("foxtrot",done[5])
	for _,d:=range done{
		<-d
	}
}


// func slowgreet(phrase string,done chan bool){
// 	time.Sleep(2*time.Second)
// 	fmt.Println("hello!",phrase)
// 	done <- true
// 	close(done)
// }

// func main(){
// 	done:=make(chan bool)
// 	go greet("alpha",done)
// 	go greet("beta",done)
// 	go slowgreet("charlie",done)
// 	go greet("delta",done)
// 	go slowgreet("echo",done)
// 	go slowgreet("foxtrot",done)
// 	for range done{

// 	}
// }


//select -manage channels
/*select{
case err:=<-channel:

case <-done:
}*/
//main idea for select is to select different cases for different channels

//defer keyword called once the surrounding method or function exucted by any of cases
//defer file.Close()
// defer was used in situations when ever we need some task to be done after execution of method.
