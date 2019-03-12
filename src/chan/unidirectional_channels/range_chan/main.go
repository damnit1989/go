package main

import "fmt"

func info(cn chan int){
	for i:= 0; i <= 10; i++{
		cn<-i
	}
	close(cn)
}

func main(){
	chl := make(chan int)
	go info(chl)
	for v := range chl{
		//fmt.Println(k)
		fmt.Println(v)
	}
	fmt.Println("this is deadlock")
}