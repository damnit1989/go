package main

import (
	"fmt"
	"math"
)

type circle struct{
	r float64
}

type shape interface{
	area() float64
}

func (c circle) area() float64{
	return math.Pi * c.r * c.r
}

func info(s shape){
	fmt.Println(s.area())
}

func main(){
	c := circle{5}
	info(&c)
}
