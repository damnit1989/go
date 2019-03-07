package main

import "fmt"

func Hello(ch chan string){
	fmt.Printf("hello function finish\n")
	ch <- "i am  comming"
	fmt.Printf("hello  cant't bloack\n")
	fmt.Printf("block hello\n")
}

func main(){
	c := make(chan string)
	go Hello(c)
	info := <- c
	fmt.Println(info)
	fmt.Printf("main gorouting finish\n")
}