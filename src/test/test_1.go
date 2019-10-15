//一般而言，当panic异常发生时，程序会中断运行，并立即执行在该goroutine（可以先理解成线程，在第8章会详细介绍）中被延迟的函数（defer 机制）。随后，程序崩溃并输出日志信息。日志信息包括panic value和函数调用的堆栈跟踪信息

package main

//import "os"
import "fmt"

func main(){
	defer func(){
	    err := recover()
		//fmt.Println(err)
		if err != nil{
			fmt.Println(err)
			fmt.Printf("%T\n", err)
		}
	    
  	    fmt.Println(1, 2, 3)
	}()
	//os.Exit(0)
	panic("wrong")
	panic("wrong two")
}
