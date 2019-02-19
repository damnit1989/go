package main

import "fmt"

func main() {
	customerNumber := make([]int, 3)
	customerNumber[0] = 7	
	customerNumber[1] = 10
	customerNumber[2] = 15
	//customerNumber[3] = 15

	fmt.Println(customerNumber[0])
	fmt.Println(customerNumber[1])
	fmt.Println(customerNumber[2])

	greeting := make([]string, 3, 5)
	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "dias!"
	//greeting[3] = "dias!"
	
	fmt.Println(greeting[2])
	fmt.Println(greeting[0])
}