package main

import(
	"fmt"
	"sort"
)

func main(){
	s := []string{"Zeno", "John", "Al", "Jenny"}
	sort.Strings(s)
	fmt.Println(s)

	/*cannot use s (type []string) as type sort.Interface in argument to sort.Sort:
	[]string does not implement sort.Interface (missing Len method)*/	
	
	//sort.Sort(s)
	//fmt.Println(s)
}