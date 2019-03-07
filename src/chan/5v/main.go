package main

import "fmt"

func area1(i int, c chan int){
	c<- i*i*i
}

func area2(i int, c chan int){
	c<-i+i+i
}

func main(){
	i := 10
	
	c1 := make(chan int)
	c2 := make(chan int)
	
	go area1(i, c1)
	
	go area2(i, c2)
	
	info1 := <-c1
	info2 := <-c2
	fmt.Println(info1 + info2)
}