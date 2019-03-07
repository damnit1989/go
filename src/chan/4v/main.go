package main

import(
	"fmt"
	"time"
)

func hello(done chan bool){
	<-done
	//time.Sleep(4 * time.Second)
	fmt.Println("hello go routing is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello")

}

func main(){
	done := make(chan bool)
	fmt.Println("Main going to call hello go gorouting")
	go hello(done)

	done <- true
	
	fmt.Println("Main received data")
}