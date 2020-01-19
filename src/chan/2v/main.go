package main

import(
	"fmt"
	"time"
)

func hello(){
	fmt.Println("Hello world gorouting")
}

func hello1(){
	time.Sleep(1 * time.Second)
	fmt.Println("Hello1 world goruting")
}
func main(){
	go hello1()
	time.Sleep(1 * time.Second)
	fmt.Println("main function")
}
