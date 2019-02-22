//变量名是给编译器看的，编译器根据变量是局部还是全局分配内存地址或栈空间，所谓的变量名在内存中不存在，操作时转换成地址数存放在寄存器中了
package main

import "fmt"

func main() {
	var a int = 20
	var ip *int
	ip = &a
	
	fmt.Printf("a 变量的地址是：%x\n",&a)
	fmt.Printf("ip 变量的存储地址:%x\n",ip)
	fmt.Printf("ip 变量自身的存储地址:%x\n",&ip)
	fmt.Printf("*ip 变量的值：%d\n", *ip)
	
	type data struct{
		a int
	}
	
	var d = data{1234}
	var p *data
	
	p = &d
	fmt.Printf("%p, %v\n", p, p.a)
	
	//空指针
	
	var myPoint *int
	fmt.Println(myPoint)
	
}