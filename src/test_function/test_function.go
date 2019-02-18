package main

import "fmt"

type VarStr struct{
	Name string
	age int
}

func (v VarStr) setName(name string){
	v.Name = name
}

func (v VarStr) setAge(age int){
	v.age = age
}

func (v VarStr) getName() string {
	fmt.Println(v.Name)
	return v.Name
}

func (vs VarStr) getAge() int {
	fmt.Println(vs.age)
	return vs.age
}

func returnStruct() VarStr {
	return VarStr{"aa",123}
}

func main(){
	v := VarStr{"lmm",20}
	v.getName()
	v.getAge()
	fmt.Println(v)
	
	v.setName("aaa")
	v.setAge(22)
	v.getName()
	v.getAge()
	fmt.Println(v)
	fmt.Println(returnStruct())
}