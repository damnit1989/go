package main

import "fmt"

func main(){
	m := map[int]string{
		0: "aaa",
		1: "bbb",
		2: "ccc",
		3: "ddd",
	}
	
	for k, v := range m{
		fmt.Println("k=",k, "v=", v)
	}
}