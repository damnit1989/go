package main

import "fmt"

func multiple(x,y string)(string,string){
	return y,x
}

func main(){
	//var x ,y string;
	x,y := multiple("123","456")
	fmt.Println("the resultes is",x,y);
}
