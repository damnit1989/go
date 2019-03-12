//我们知道select语句将被阻塞，直到执行其中一个case。
//在这种情况下，select语句没有任何case，因此它将永远阻塞导致死锁。

package main

func main(){
	select{
	}
}