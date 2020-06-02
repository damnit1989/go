package utils

import (
	"log"
	"os"
	"runtime"
)

// import "os"
// var files map[string]int64
// var files = make(map[string]int64, 10)
var files = make(map[string]int64, 0)

func findFiles(filename string) map[string]int64 {
	// var files []string
	pfile, err := os.Open(filename)
	if err != nil {
		log.Panic(err)
	}

	fileInfo, err := pfile.Readdir(0)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range fileInfo {
		if v.IsDir() {
			findFiles(filename + "/" + v.Name())
		} else {

			files[v.Name()] = v.Size()
			// files = append(files, v.Name())
		}
	}
	return files
}

func test() {
	runtime.Version()
}
