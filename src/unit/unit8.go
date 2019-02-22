package main

import "fmt"

func main(){
	s := "Go编程"
	fmt.Println(len(s))
	//m := len(string(rune("编")))
	m := len("编")
	fmt.Println(m)
}