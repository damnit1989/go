//多路复用
package main

import(
	"fmt"
	"time"
)

func worker_one(chl chan string){
	time.Sleep(6 * time.Second)
	chl <- "worker one"
	//close(chl)
}

func worker_two(chl chan string){
	time.Sleep(3 * time.Second)
	chl <- "worker two"
	//close(chl)
}


func main(){
	chan1 := make(chan string)
	chan2 := make(chan string)
	
	go worker_one(chan1)
	go worker_two(chan2)

	select{
		case data := <- chan1:
			fmt.Println(data)
		case data := <- chan2:
			fmt.Println(data)
		//default:
			//fmt.Println("123456")
	}

	
	fmt.Printf("all done\n")
}
