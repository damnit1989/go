package main

import "fmt"

func main(){
	buf_chl := make(chan string, 2)
	
	buf_chl <- "abc"
	fmt.Println(<-buf_chl)
	
	buf_chl <- "def"
	buf_chl <- "hij"
	
	fmt.Println(<-buf_chl)
	fmt.Println(<-buf_chl)
}