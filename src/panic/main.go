//请注意，与使用异常处理许多错误的某些语言不同，在Go中，尽可能使用错误指示返回值是惯用的
package main

//import "os"
import "log"
import "fmt"

func main(){
	//panic("a problem")
	log.Fatal("a problem")
	//_, err := os.Create("/tmp/file")
    //if err != nil {
    //    panic(err)
    //}
	fmt.Println("end")
}