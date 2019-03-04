package main

import (
	//"errors"
	"fmt"
)

type MyError struct{
	s string
}

func New(msg string) error{
	return MyError{msg}
}

func (e MyError) Error() string{
	return fmt.Sprintf("v: %v", e.s)
}

func main(){
	msg := New("xxxxxxxx")
	if msg != nil{
		fmt.Println(msg)
	}
}

