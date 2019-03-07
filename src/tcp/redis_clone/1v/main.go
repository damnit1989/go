package main

import (
	"fmt"
	"net"
	//"log"
	"bufio"
)

func handle(conn net.Conn){
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		info := scanner.Text()
		fmt.Println(info)
	}
}

func main(){
	ln, err := net.Listen("tcp", ":")
	if err != nil{
		panic(err)
	}	
	//defer ln.Close() 
	//defer ln.Close()
	for{
		conn, err := ln.Accept()
		if err != nil{
			panic(err)
		}
		handle(conn)
	}
}