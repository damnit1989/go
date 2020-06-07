package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"sort"
	// "time"
)

func main() {

	var arr = []string{"wne", "xwo"}
	sort.Strings(arr)
	// fmt.Printf("%T", arr)
	pos := sort.SearchStrings(arr, "xwo")
	fmt.Println(pos)

	// var str = "123456"

	h := md5.New()
	io.WriteString(h, "123456")
	fmt.Printf("%x", h.Sum(nil))
}
