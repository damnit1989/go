package main

import "fmt"

//普通变量参数
func swap(x, y int){
	tmp := x
	x = y
	y = tmp
}

//指针参数传值
func pSwap(x, y *int){
	tmp := *x
	*x = *y
	*y = tmp
}

func main() {
	var x, y int = 3, 4
	swap(x, y)
	fmt.Println(x, y)
	
	//引用地址
	pSwap(&x, &y)
	fmt.Println(x, y)
}
