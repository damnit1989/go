package main

import "fmt"

//固有接口
type Before interface{
	Attack()
	Defence()
}

type ForeignPerson struct{
	Name string
}

func (this *ForeignPerson) Attack() {
	fmt.Printf("attack")
}

func (this *ForeignPerson) Defence() {
	fmt.Printf("defence")
}

type ChinaPerson struct {
	Name string
}



func main() {
	p1 := ForeignPerson{"james"}
	p2 := ForeignPerson{"kebi"}
	
	ym := ChinaPerson{"姚明"}
}
