package main

import "fmt"

func main(){
	m := map[string]string{}
	fmt.Println(m)
	
	m["a"] = "abc"
	m["b"] = "123"
	fmt.Println(m)
}