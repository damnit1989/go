package main

import (
	"fmt"
	"log"
	"net"
)

func main(){
	conn, err := net.Dial("tcp", "192.168.17.134:12345")
	if err != nil {
		log.Fatal(err)
		return 
	}

	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		fmt.Printf("请输入发送的内容:")
		fmt.Scan(&buf)
		fmt.Printf("发送的内容:%s\n", string(buf))
		
		conn.Write(buf)
		n, err := conn.Read(buf)
		if err != nil{
			fmt.Println(err)
			return 
		}

		result := buf[:n]
		fmt.Printf("接收的数据[%d]:%s\n", n, string(result))
	}
}