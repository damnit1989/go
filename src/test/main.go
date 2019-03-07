package main

import(
	"fmt"
	"strings"
)

func main(){
	str := "hello world i like,you"
	ret := strings.Fields(str)
	fmt.Println(ret[0])
	fmt.Println(ret[1])
	fmt.Println(ret)
	for key, val :=  range ret{
	 	fmt.Println(key, "--", val)
	}
}
