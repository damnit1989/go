// 常量模拟枚举
package main

import "fmt"

const PI = 3.14
const (
	FlagNone = 1 << iota
	FlagRed
	FlagGreen
	FlagBLue
)

type alse int


func (r alse) info(){
	println("dddddd")
	// fmt.Println("dddddd")
}

func (r alse) String(){
	// return "ssss"
	println("ssssssss")
}
// cannot define new methods on non-local type int
// func (s int) info(){}

var a alse = 3

func main(){
	fmt.Println(PI)
	fmt.Println(FlagNone,FlagRed,FlagGreen,FlagBLue)
	fmt.Printf("%b,%b,%b,%b\n",FlagNone,FlagRed,FlagGreen,FlagBLue)
	a.info()
	a.String()
	fmt.Printf("%d",a)
}