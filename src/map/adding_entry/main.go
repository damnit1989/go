package main

import "fmt"

func main(){
	m := map[string]string{
		"a":"aa",
		"b":"bb",
	}
	
	m["c"] = "cc"
	
	fmt.Println(m)
}