package main

import "fmt"

func main(){
	var m map[string] string;
	
	fmt.Println(m)
	fmt.Println(m == nil)

	n :=  map[string]string{}	
	fmt.Println(n)
	fmt.Println(n == nil)
	if len(n) == 0  {
	    fmt.Printf("the n is empty map")
	}
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
