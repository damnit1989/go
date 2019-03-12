// 缓冲信道的容量是信道可以容纳的值的数量。这是我们使用make函数创建缓冲通道时指定的值。
// 缓冲通道的长度是当前在其中排队的元素数

package main

import "fmt"

func main(){
	ch := make(chan string, 3)
	ch <- "abc"
	ch <- "def"
	
	fmt.Println("capacity is", cap(ch))
	fmt.Println("length is", len(ch))
	fmt.Println("read value", <-ch)
	fmt.Println("new length is", len(ch))
	fmt.Println("new capacity is", cap(ch))
}