package main

import "fmt"

func main(){
	m := map[string]string{
		"a":"aaa",
		"b":"bbb",
	}
	
	m["c"] = "ccc"
	fmt.Println(m)
	
	m["c"] = "ddd"
	fmt.Println(m)
}