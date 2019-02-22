package main

import(
	"fmt"
	"encoding/json"
)

type person struct{
	First string
	Last string 
	Age int `json:"world"`
}

func main(){
	var p person
	bs := []byte(`{"First":"lmm", "Last":"mm", "world":20}`)
	err := json.Unmarshal(bs, &p)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(p)
	fmt.Println(p.First)
	fmt.Println(p.Last)
	//fmt.Println(p.world)
	fmt.Println(p.Age)
}