package main

import "fmt"

type animal struct{
	sound string
}

type dog struct{
	animal
	friendly bool
}

type cat struct{
	animal
	annoying bool
}

/*****当参数类型为接口类型时，则可传任何类型的参数*****/

//第一种方式定义接口参数类型
func specs(a interface{}){
	fmt.Printf("类型:%T\n", a)
	fmt.Println(a)
}

//第二种方式定义接口参数类型
//type inter interface{}
//func specs(a inter){
//	fmt.Println(a)
//}

func main(){
	fido := dog{animal{"woof"}, true}
	fifi := cat{animal{"meow"}, true}
	a := 10
	
	specs(fido)
	specs(fifi)
	specs(a)
}