// 结构体内嵌
package main

import "fmt"

// 普通类型嵌套
type  Normal struct{
	int 
	string 
	bool
}

// 结构体类型嵌套
type Nest struct{
	Normal
}

func main(){
	n := Normal{1,"aa",true}
	fmt.Println(n)

	var nest Nest
	nest.int = 1
	nest.string = "aa"
	nest.bool = true
	fmt.Println(nest)
}

