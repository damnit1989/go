package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup){
	fmt.Println("started Gorouting", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	
	//计数器减一
	wg.Done()
}

func main(){
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++{
		wg.Add(1)
		go process(i, &wg)
	}
	
	//等待计数器变为零
	//那就是wg.Done（）被调用三次，计数器将变为零，主要的Goroutine将被解锁
	wg.Wait()
	fmt.Println("All go routines finished executing")
}