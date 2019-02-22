package main

import "fmt"

const MAX = 3

func main() {
	a := [MAX]int{10,100,200}
	for i:=0; i< MAX; i++{
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}
	
	fmt.Printf("============================\n")
	
	b := []int{9, 99, 999}
	var ptr [MAX]*int
	for j := 0; j<MAX;j++{
		ptr[j] = &b[j]
	}
	
	for j := 0; j<MAX;j++{
		fmt.Printf("a[%d] = %d\n", j, *ptr[j])
	}
}