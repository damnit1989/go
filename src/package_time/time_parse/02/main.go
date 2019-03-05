package main

import(
	//"fmt"
	"time"
	"log"
)

func main(){
    timeAsString := "01/22/2012"
    //fmt.Println(time.Parse("01/01_this-does-not-compile/2016", timeAsString))
    _, err := time.Parse("01/01_this-does-not-compile/2016", timeAsString)
    if err != nil{
	log.Fatal(err)
    }else{
    	log.Fatal("sdfsfsdf")
    }
}

