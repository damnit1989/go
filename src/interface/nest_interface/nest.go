package main

import "io"

type device struct{

}

func(d *device) Write(b []byte)(n int,err error){
	return 0,nil
}
func(d *device) Close() error{
	return nil
}

func main(){
	var wl io.WriteCloser = new(device)
	wl.Write(nil)
	wl.Close()
}