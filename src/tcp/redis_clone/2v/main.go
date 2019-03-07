package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
	"fmt"
)

func handle(conn net.Conn){
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		ln := scanner.Text()
		fs := strings.Fields(ln)
		//if len(fs) < 1{
		//	continue
		//}
		fmt.Fprintf(conn, "%s\n", fs)
		
		switch fs[0]{
			case "GET":
				io.WriteString(conn,"COMMAND GET\n")
			case "SET":
				io.WriteString(conn,"COMMAND SET\n")
			case "DEL":
				io.WriteString(conn,"COMMAND DEL\n")
			default:
				io.WriteString(conn,"INVALID COMMAND "+fs[0]+"\n")
		}
	}
}

func main(){
	li, err := net.Listen("tcp", ":8642")
	if err != nil{
		log.Fatal(err)
	}
	defer li.Close()
	
	for{
		conn, err := li.Accept()
		if err != nil{
			log.Fatalln(err)
		}
		handle(conn)
	}
}