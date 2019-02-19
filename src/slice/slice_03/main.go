package main

import "fmt"

func main() {
	greeting := []string{
		"Good morning!",
		"Bonjour!",
		"dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}

	for i, currentEntry := range greeting {
		fmt.Println(i, currentEntry)
	}

	for j := 0; j < len(greeting); j++ {
		fmt.Println(greeting[j])
	}

	testSlice()
}

func testSlice() {
	var results []int
	fmt.Println(results)

	mySlice := []string{"a","b","c","g","m","z"}
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4])  // slicing a slice
	fmt.Println(mySlice[2])    // index access; accessing by index
	fmt.Println("myString"[2]) // index access; accessing by index
	fmt.Println(mySlice[:])	
}