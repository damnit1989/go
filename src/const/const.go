package main

import "fmt"

const (
	name = "lmm"
	age
	aa
	bb
)

//枚举 类型
const (
	START = iota + 1
	END
	PREPARE
)

const (
	OPEN = iota
	CLOSE
)

func main(){
	fmt.Println(name, age, aa, bb)
	fmt.Println(START, END, PREPARE)
	fmt.Println(OPEN, CLOSE)
}