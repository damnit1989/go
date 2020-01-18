package main

import (
	"fmt"
	"time"
)

func sumLeft(c chan int) int {
	sum := 0
	for i:= 51;i<=100;i++{
		sum += i
	}
	c <- sum
	// fmt.Println(sum)
	return sum
}

func sumRight(c chan int) int {
	sum := 0
	for i:= 1; i<= 50;i++{
		sum += i 
	}
	// fmt.Println(sum)
	// c <- sum
	return sum
}


func main(){
	var ch1,ch2 chan int
	ch1 = make(chan int)
	ch2 = make(chan int)

	// sum := 0

	// go sumLeft(ch1)
	// go sumRight(ch2)

	go func(){
		sum_left := 0
		for i:= 1;i <= 50;i++{
			sum_left += i
			// fmt.Println(sum)
		}

		ch1 <-sum_left
	}()
	go func(){
		sum_right := 0
		for i :=51;i<= 100;i++{
			sum_right += i
		}
		ch2 <- sum_right
	}()

	
	sum_left := <- ch1
	sum_right := <- ch2
	time.Sleep(1 * time.Second)
	fmt.Println(sum_left,sum_right)
	fmt.Println(sum_left + sum_right)
	fmt.Println(&ch1)
}