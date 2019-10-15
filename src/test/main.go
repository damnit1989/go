package main

import(
	"fmt"
	"strings"
	"time"
)

func test(){
	panic("errors")
}

func main(){

	ch := make(chan string)
	go func(){
	    ch <- "lmm"
	    fmt.Println("this is go func")
	}()
	
	str := "hello world i like,you"
	ret := strings.Fields(str)
	fmt.Println(ret[0])
	fmt.Println(ret[1])
	fmt.Println(ret)
	for key, val :=  range ret{
	 	fmt.Println(key, "--", val)
	}
	
	fmt.Println(time.Now())
//	fmt.Println(time.Date())
	//test()
	
	//time.Sleep(time.Second)

	//<- ch
}
