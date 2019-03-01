package main

import (
	"fmt"
	//"runtime"
)

/*func init() {
    numcpu := runtime.NumCPU()
    runtime.GOMAXPROCS(numcpu) // 尝试使用所有可用的CPU
}*/

func main() {
	go foo()
	go bar()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for j := 0; j < 45; j++ {
		fmt.Println("Bar:", j)
	}
}