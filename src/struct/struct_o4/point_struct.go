package main

import (
	"fmt"
)

type persion struct{
	name string
	age int
}

func main(){

	// 普通类型结构体
	p := persion{"lmm",23}
	
	fmt.Println(p)
	fmt.Println(p.name,p.age)
	fmt.Printf("%T\n",p)

	// 指针类型结构体
	// pp := new(persion{"lmm",23})
	pp := &persion{"lmm",23}
	fmt.Println(pp)
	fmt.Println(pp.name,pp.age)
	fmt.Printf("%T\n",pp)

	p3 := new(persion)
	p3.name = "lmm"
	p3.age = 23
	fmt.Println(p3)
	fmt.Println(p3.name,p3.age)
	fmt.Printf("%T",p3)

}