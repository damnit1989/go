package main

import "fmt"

type vehicles interface{}

type vehicle struct{
	Seats int
	MaxSpeed int
	Color string
}

type car struct{
	vehicle
	Wheels int
	Doors int
}

type plane struct{
	vehicle
	Jet bool
}

type boat struct{
	vehicle
	Length int
}

func main(){
	prius := car{}
	tacoma := car{}
	bmw528 := car{}
	boeing747 := plane{}
	boeing757 := plane{}
	boeing767 := plane{}
	sanger := boat{}
	nautique := boat{}
	malibu := boat{}
	
	//也可以添加整形数据
	a := 10
	
	//接口类型的切片，可以添加任何类型的数据
	rides := []vehicles{prius, tacoma, bmw528, boeing747, boeing757, boeing767, sanger, nautique, malibu, a}
	
	for key, value := range rides{
		fmt.Println(key, "-", value)
	}
}