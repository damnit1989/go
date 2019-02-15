package main

import(
	"fmt"
	"math/rand"
)

func main(){
	fmt.Println("my favourite number is",rand.Intn(10))
	fmt.Printf("my favourite number is %d",rand.Intn(10))
}