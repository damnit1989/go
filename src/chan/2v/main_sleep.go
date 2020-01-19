package main

import(
	"fmt"
	"time"
)

func hello(){
	time.Sleep(1 * time.Second)
	fmt.Println("Hello world gorouting")
}

func main(){
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("main function")
}
