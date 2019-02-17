package main

import(
	"fmt"
	"encoding/json"
	"os"
	"bytes"
	"log"
)

func test(){
	type  ColorGroup struct{
		ID int
		Name string
		Colors []string
	}
	group := ColorGroup{
		ID: 1,
		Name: "Reds",
		Colors: []string{"crimson","red","Ruby","maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil{
		fmt.Println("error:",err)
	}
	//fmt.Println(b)
	os.Stdout.Write(b)
}

func main(){
	type Road struct{
		Name string
		Number int
	}
	
	roads := []Road{{"Diamond Fork",29},{"Sheep Creak",51}}
	
	b,err := json.Marshal(roads)
	if err != nil{
		log.Fatal(err)
	}

	var out bytes.Buffer
	json.Indent(&out,b,"=","\t")
	out.WriteTo(os.Stdout)
	fmt.Printf("\n")
	
	test()
}