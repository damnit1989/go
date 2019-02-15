package  main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func addParam(x,y int) int {
	return x + y
}

func main(){
	fmt.Println(add(10.0,11))
	fmt.Println(addParam(10.0,11))
}