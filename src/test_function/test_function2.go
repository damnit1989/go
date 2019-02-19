package main

import "fmt"
//import "net/http"

type Persion struct{
	Name string
	Age int
}

func (p Persion) getName() string {
	return p.Name
}

func (p *Persion) getNameByPointer() string{
	return p.Name
}

func main(){
	p := Persion{"lmm", 20}
	fmt.Println(p.getName())
	fmt.Println((&p).getNameByPointer())
	fmt.Println(p.getNameByPointer())
	
}