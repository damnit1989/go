package main

import (
	"fmt"
	"encoding/json"
)

type person struct {
	First string
	Last  string `json:"-"`//貌似转json的时候会忽略掉该字段
	Age   int	 `json:"wisdom score"`//json键重命名
}

func main() {
	p1 := person{"James", "Bond", 20}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Println(string(bs))
}