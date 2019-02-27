package main

import "fmt"

func main(){
    if true || false {
	fmt.Printf("this run\n")
    }
    
    //不允许使用or连接
    if true or false{
	fmt.Printf("this why\n")
    }
}
