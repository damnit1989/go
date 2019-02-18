package main

func test(){
	
	//第一种方式
	defer func(){
		if err := recover(); err != nil{
			println(err.(string))
		}
	}()
	
	//第二种方式
	/*f := func(){
		if err := recover(); err != nil{
			println(err.(string))
		}
	}
	defer f()*/
	
	panic("something wrong")
}

func main(){
	test()
}