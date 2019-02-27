package main

import "fmt"

func main(){
	m := map[string]string{
		"a" : "aaa",
		"b" : "bbb",
		"c" : "ccc",
	}
	
	fmt.Println(m)
	//删除键值
	delete(m,"a")
	fmt.Println(m)
	
}
