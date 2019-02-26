package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) getR() float64{
	return c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
	//fmt.Println("R", s.getR())
}

func main() {
	c := circle{5}
	
	//运行出错
	info(c)
}