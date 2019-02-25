package main

import(
	"fmt"
	"sort"
)

func main(){
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println(n)
	
	//升序排序
	//sort.Sort(sort.IntSlice(n))
	
	//降序排序
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	
	//sort.Reverse(sort.IntSlice(n))
	fmt.Println(n)
}