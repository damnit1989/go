//defer 栈
package main

import(
	"fmt"
)

func main(){
	defer fmt.Printf("world\n")
	for i := 0;i<10;i++{
		defer fmt.Println(i)
	}
	fmt.Printf("hello\n")
}