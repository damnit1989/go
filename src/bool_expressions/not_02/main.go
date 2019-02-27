package main

import "fmt"

func main(){
    if !true{
	fmt.Printf("this not run\n")
    }

    if !false{
	fmt.Printf("this run\n")
    }
}
