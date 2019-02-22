package main

import (
	"fmt"
	"encoding/json"
)

type person struct{
	Name string
	Age int
	notExported  int
}

//为什么结构体中字段名是小写的无法转成json
type personStr struct{
	first  string
	last string
	ege int
}

func main() {
	p := person{"lmm",12, 42}
	
	data, _ := json.Marshal(p)
	fmt.Println(p)
	fmt.Println(data)
	fmt.Printf("%T\n", data)
	fmt.Println(string(data))
	
	p1 := personStr{"jams", "Bond", 20}
	fmt.Println(p1)
	bs, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Println(string(bs))
	
}