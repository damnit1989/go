package main

func info(a []int,f func(int)){
	for _,v := range a{
		f(v)
	}
}
func callback(p int){
	println(p)
}

// var arr = [5]int{1,2,3,4,5}

func main(){
	p := []int{1,2,3,4,5}
	// info(p,func(a int ){
	// 	println(a)
	// })
	
	info(p,callback)
}