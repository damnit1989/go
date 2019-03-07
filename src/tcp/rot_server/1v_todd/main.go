package main

import (
	"fmt"
	"bufio"
	"net"
	"strings"
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
		
		scanner := bufio.NewScanner(conn)
		for scanner.Scan(){
			line := strings.ToLower(scanner.Text())
			bs := []byte(line)
			var rot13 = make([]byte, len(bs))
			for k, v := range bs{
				if v <= 109{
					rot13[k] = v + 13
				}else{
					rot13[k] = v - 26 + 13
				}
			}
			fmt.Fprintf(conn, "bs: %s\n", bs)
			fmt.Fprintf(conn, "rot13: %s\n", rot13)
		}
	}
}