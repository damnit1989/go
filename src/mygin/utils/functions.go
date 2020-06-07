package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

// FileMd5 return the file md5
func FileMd5(filename string) string{
	reader, err := os.Open(filename)
	if err != nil {
		log.Panic("文件读取失败")
	}

	h := md5.New()
	if _,err := io.Copy(h,reader);err != nil {
		log.Panic("文件内容复制错误")
	}

	return fmt.Sprintf("%x",h.Sum(nil))
}