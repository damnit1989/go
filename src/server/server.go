package main

import(
	"fmt"
	"log"
	"net"
	"strings"
)

func dealConn(conn net.Conn){
	defer conn.Close()

	ipAddr := conn.RemoteAddr().String()
	fmt.Println(ipAddr, "连接成功")

	buf := make([]byte, 1024)
	for{
		n, err := conn.Read(buf)
		if err != nil{
			fmt.Println(err)
			return 
		}

		result := buf[:n]
		fmt.Printf("接收到数据来自[%s] ==> [%d]:%s\n",ipAddr,n,string(result))
		if "exit" == string(result){
			fmt.Println(ipAddr,"退出连接")
			return 
		}
		conn.Write([]byte(strings.ToUpper(string(result))))
	}
}

func main(){
	listenner, err := net.Listen("tcp", "192.168.17.134:12345")
	if err != nil{
		log.Fatal(err)
	}

	for{
		conn, err := listenner.Accept()
		if err != nil{
			log.Println(err)
			continue
		}
		go dealConn(conn)
	}
}