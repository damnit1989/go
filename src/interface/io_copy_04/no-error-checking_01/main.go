package main

import(
	"bytes"
	"io"
	"net/http"
	"os"
	"strings"
)

func main(){
	msg := "Do not dwell in the past"
	rdr := strings.NewReader(msg)
	io.Copy(os.Stdout, rdr)
	
	rdr2 := bytes.NewBuffer([]byte(msg))
	io.Copy(os.Stdout, rdr2)
	
	res, _ := http.Get("http://www.baidu.com")
	io.Copy(os.Stdout, res.Body)
	res.Body.Close()
}