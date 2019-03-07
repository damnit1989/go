package main

import "fmt"

func sendData(sendch chan<- int){
	sendch <- 10
}

func main(){
	chnl := make(chan int)
	
	//单独线程处理
	go sendData(chnl)
	fmt.Println(<-chnl)
}
