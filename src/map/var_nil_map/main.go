package main

import "fmt"

func main(){
	var m map[string] string;
	
	fmt.Println(m)
	fmt.Println(m == nil)
	
	//if m == "" || m == 0{
	//    fmt.Println("m is what")
	//}
	//fmt.Printf("nil: %T", nil)
	
	//if nil == "" || nil == 0{
	//	fmt.Printf("nil is what !!!")
	//}
	
	// add these lines:
	//panic: assignment to entry in nil map	

	//m["Tim"] = "Good morning"
	//m["Jenny"] = "Bonjour"
}
