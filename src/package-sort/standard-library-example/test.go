package main

import "fmt"

type person struct{
	Name string
	Age int
}

//实现string()方法
func (p person) String() string{
	return fmt.Sprintf("My Name is %s,Age is %d", p.Name, p.Age)
}

func main(){
	p := person{"xxx", 12}
	fmt.Println(p)
}