package main

import(
	"fmt"
	"time"
)

func server1(ch chan string){
	ch <- "from server1"
}

func server2(ch chan string){
	ch <- "from server2"
}

func main(){
	chan1 := make(chan string)
	chan2 := make(chan string)
	
	go server1(chan1)
	go server2(chan2)
	
	time.Sleep(2*time.Second)
	
	select{
		case ret := <- chan1:
			fmt.Println(ret)
		case ret := <- chan2:
			fmt.Println(ret)
	}
}