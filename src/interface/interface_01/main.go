package main

import "fmt"

type square struct{
	side float64
}

func (z square) area() float64{
	return z.side * z.side
}

type rectangle struct{
	LongSide float64
	ShortSide float64
}

func (z rectangle) area() float64{
	return z.LongSide * z.ShortSide
}


type shape interface{
	area() float64
}

func info(z shape){
	fmt.Println(z)
	fmt.Println(z.area())
}

func main(){
	s := square{10}
	rtg := rectangle{20, 10}
	//fmt.Printf("%T\n", s)
	
	info(s)
	info(rtg)
}
