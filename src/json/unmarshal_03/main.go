package main

import (
	"fmt"
	"encoding/json"
)

func main(){
	var jsonBlob = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	]`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	
	fmt.Printf("%+v\n", animals)
	fmt.Printf("%+v\n", animals[0])
	fmt.Printf("%+v\n", animals[1])
	for i := 0;i < len(animals);i++{
		fmt.Println(animals[i].Name,animals[i].Order)
	}
	
	
}