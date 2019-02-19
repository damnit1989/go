package main

import "fmt"

func main() {
	mySlice := []int{1,3,5,7,9,11}
	//mySlice := [1,3,5,7,9,11]
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)

	xs := []int{1,3,5,7,9,11}
	for i, value := range xs {
		fmt.Println(i, "-", value)
	}

	//切片
	fmt.Println(xs[1:])
	fmt.Println(xs[1:5])
	fmt.Println(xs[:5])
}