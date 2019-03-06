package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main(){
	conn, err := net.Dial("tcp", "192.168.17.134:8642")
	if err != nil{
		panic(err)
	}
	defer conn.Close()
	
	bs, _ := ioutil.ReadAll(conn)
	fmt.Println(string(bs))
}