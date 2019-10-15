package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("%T", os.Environ())
	//fmt.Println(os.Environ())
	for _, e := range os.Environ() {
		env := strings.Split(e, "=")
		fmt.Println(env[0])
		fmt.Println(env[1])
	}
}