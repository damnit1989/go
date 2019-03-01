package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main(){
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil{
		log.Fatal(err)
	}
	
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	
	scanner.Split(bufio.ScanWords)
	buckets := make([]int, 12)
	for scanner.Scan(){
		n := hashBucket(scanner.Text(), 12)
		buckets[n]++
	}
	fmt.Println(buckets)
}

func hashBucket(word string, buckets int)int{
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}