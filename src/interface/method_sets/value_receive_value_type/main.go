package main

import (
	"fmt"
	"math"
)

type circle struct{
	r float64
}

func (c circle) area() float64{
	return math.Pi * c.r * c.r
}

type shape interface{
	area() float64
}

//func info(s shape){
//	fmt.Println(s.area())
//}

func info(s shape){
	fmt.Println(s.area())
}

func main(){
	c := circle{5}
	fmt.Println(c)
	//fmt.Printf("%T", c)
	info(c)
}

