package main

import "fmt"

func main(){
	m := map[int]string{
		0:"aaa",
		1:"bbb",
		2:"ccc",
		3:"ddd",
	}

	fmt.Println(m)
	
	delete(m, 2)
	if val, exists := m[2];exists{
		fmt.Println("That value exists.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)		
	}else{
		fmt.Println("That value doesn't exist.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)	
	}
	
	fmt.Println(m)
}