package main

import "fmt"

func testFunction(list []int, f func(int)int) []int{
	//for i := 0;i < len(list);i++{
	//	list[i] = f(list[i])
	//}
	
	for k, v := range list{
		list[k] = f(v)
	}
	return list
}

func main(){
	list := []int{1,2,3,4,5,6,7,8}
	new_list := testFunction(list, func(i int)int{
		return i * i
	})
	fmt.Println(new_list)
	fmt.Println(list)
}
