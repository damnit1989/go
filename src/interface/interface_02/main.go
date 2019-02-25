package main

import(
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct{
	radius float64
}

//接口类型
type shape interface{
	area() float64
}

//实现了接口类型
func (s square) area() float64{
	return  s.side * s.side
}

//实现了接口类型
func (c circle) area() float64{
	return math.Pi * c.radius * c.radius
}

//定义方法，参数类型为接口类型
func info(z shape){
	fmt.Println(z)
	fmt.Println(z.area())
}

func main(){
	s := square{10}
	c := circle{5}
	info(s)
	info(c)
}