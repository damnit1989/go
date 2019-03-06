package main

import(
	"bufio"
	"io"
	"net"
)

func main(){
    ln, err := net.Listen("tcp", ":8642")
	if err != nil{
		panic(err)
	}
	for{
		conn, err := ln.Accept()
		if err != nil{
			panic(err)
		}

		scanner := bufio.NewScanner(conn)
		for scanner.Scan(){
			io.WriteString(conn, scanner.Text())
		}
		
		conn.Close()
	}
	
	
}
