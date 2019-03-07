package main

import (
	"io"
	"net"
	"fmt"
)

func main(){
	ln, err := net.Listen("tcp", ":8642")
	if err != nil{
		panic(err)
	}
	defer ln.Close()
	
	for{
		conn, err := ln.Accept()
		if err != nil{
			panic(err)
		}
		go func(){
			fmt.Println("sdfsdf")
			io.Copy(conn, conn)
			conn.Close()
		}()
	}
}