package main

import(
	//"fmt"
	"encoding/json"
	"os"
)

type person struct{
	First string
	Last string
	Age int
}

func main(){
	p1 := person{"Jams", "Bond",20}
	json.NewEncoder(os.Stdout).Encode(p1)
}