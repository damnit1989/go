package main

import (
	"fmt"
	"time"
)

func producer(chnl chan int){
	for i:= 0; i < 10; i++{
		chnl <- i
		//time.Sleep(1 * time.Second)
	}
	close(chnl)
}

func main(){
	ch := make(chan int)
	go producer(ch)
	for{
		v, ok := <-ch
		if ok == false{
			break;
		}
		fmt.Println("Receiveds", v, ok)
		time.Sleep(1 * time.Second)
	}
}