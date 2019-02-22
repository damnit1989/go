//变量名是给编译器看的，编译器根据变量是局部还是全局分配内存地址或栈空间，所谓的变量名在内存中不存在，操作时转换成地址数存放在寄存器中了
package main

import "fmt"

func main() {
	s := []int{10, 10, 10}
	s[1] = 20

	p := make([]int, 3)
	p[0] = 1
	p[1] = 1
	p[2] = 1
	
	k := new([]int)
	//&k[1] = 10

	var m []int
	m = append(m,1)
	m = append(m,2)
	m = append(m,3)
	
	var n int = 10
	
	fmt.Println(s)
	fmt.Println(p)
	fmt.Printf("%T\n", k)
	fmt.Printf("%T\n", *k)
	fmt.Println(m)
	
	//读取变量内存地址
	fmt.Println(&n)
	
	//通过指针获取变量值
	fmt.Println(*&n)
	
	//直接打印值
	fmt.Println(n)
}