package main

import "fmt"

func main() {
	var a int
	var ptr *int
	var pptr **int
	var ppptr ***int

	//0xc00006c010
	//0xc00006e018
	//0xc00006e020
	
	fmt.Printf("a 变量的地址：%x\n", &a)
	fmt.Printf("ptr 指针变量的地址：%x\n", &ptr)
	fmt.Printf("pptr 指针变量的地址：%x\n", &pptr)
	
	a = 300
	ptr = &a
	pptr = &ptr
	ppptr = &pptr
	
	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)	
	fmt.Printf("指向指针的指针变量 **ppptr = %d\n", ***ppptr)	
	fmt.Printf("指向指针的指针变量 ptr = %x\n", ptr)
	fmt.Printf("指向指针的指针变量 pptr = %x\n", pptr)
	fmt.Printf("指向指针的指针变量 ppptr = %x\n", ppptr)

}
