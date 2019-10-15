package main

import "fmt"
import "os"

func main() {
    defer fmt.Println("one")
    defer fmt.Println("two")
    os.Exit(0)
    os.Exit(2)

}
