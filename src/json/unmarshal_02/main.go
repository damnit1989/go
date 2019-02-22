package main

import (
	"fmt"
	"encoding/json"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main(){
	var p person
	fmt.Println(p.First)
	fmt.Println(p.Last)
	fmt.Println(p.Age)
	
	bs := []byte(`{"First":"James", "Last":"Bond", "Age":20}`)
	fmt.Println(bs)
	json.Unmarshal(bs, &p)
	fmt.Println(p.First)
	fmt.Println(p.Last)
	fmt.Println(p.Age)
	fmt.Println(p)
	fmt.Printf("%T\n", p)
}