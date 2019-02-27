package main

import "fmt"

func main(){
    if true && false{
        fmt.Printf("this not run")
    }
}
