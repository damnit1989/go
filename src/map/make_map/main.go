package main

import "fmt"

func main(){

	//key value 都为string类型
	m := make(map[string]string)
	m["aa"] = "123"
	m["bb"] = "456"
	
	fmt.Println(m)
	fmt.Println(m["bb"])
	
	//key为string类型，value为任意类型
	m2 := make(map[string]interface{})
	m2["x"] = []int{1,2,3}
	m2["y"] = 456
	
	fmt.Println(m2)
	fmt.Println(m2["x"])
}