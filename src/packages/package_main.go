package main

import (
	"fmt"
	"package_one"
)

func test(p package_one.PackageOne){
	fmt.Println(p)
}

func main(){
	
	fmt.Println(package_one.A)

	//打印常量值
	fmt.Println(package_one.LENGHT, package_one.WEIGHT)
	
	//结构体
	pg := package_one.PackageOne{"lmm", 23}
	
	//调用结构体方法
	fmt.Println(pg)
	fmt.Println(pg.GetName())
	fmt.Println(pg.GetAge())
	
	//调用普通函数
	fmt.Println(package_one.Info())
	
	package_one.Main()
	
	test(pg)	
}
