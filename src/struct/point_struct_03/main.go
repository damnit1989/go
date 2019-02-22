package main

import (
	"fmt"
	"encoding/json"
	"log"
)

type person struct{
	Name  string
	Age  int
}

func main() {
	p_str := person{Name:"lmm", Age : 29}
	p := &p_str
	
	fmt.Println(p_str)
	fmt.Println(&p_str)
	fmt.Println(p)
	fmt.Println(*p)
	
	/*对于一些复杂类型的指针， 如果要访问成员变量的话，需要写成类似(*p).field的形式，Go提供了隐式解引用特性，我们只需要p.field即可访问相应的成员。*/
	fmt.Println(p.Name, p.Age)
	fmt.Println((*p).Name, (*p).Age)
	fmt.Println(p_str.Name, p_str.Age)
	str,err := json.Marshal(p_str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("str 的类型：%T, str的值：%s", str, string(str))
}