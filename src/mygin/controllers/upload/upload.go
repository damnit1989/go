package upload

import (
	"io/ioutil"
	"log"
	"mygin/utils"
	"os"

	"github.com/gin-gonic/gin"
)

// const UPLOAD_PATH = "../../uploads/"

var dst = "./uploads/"

func UploadPic(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(200, gin.H{
				"meg": err,
			})
		}
	}()

	file, err := c.FormFile("file")
	if err != nil {
		log.Panicln(err)
	}

	uploadConfig := map[string]interface{}{
		"FileType": []string{"application/zip", "application/rar"},
		"FileSzie": 1024 * 1024 * 2,
	}
	upload := utils.NewUpload(uploadConfig)
	fileInfo := upload.DoUpload(file)
	c.JSON(200, fileInfo)
}

func MultUpload(c *gin.Context) {
	multForm, err := c.MultipartForm()
	if err != nil {
		log.Panic(err)
	}
	files := multForm.File["upload[]"]
	for _, v := range files {
		c.SaveUploadedFile(v, dst)
	}
}

// 断点续传
func BreakPointUpload(c *gin.Context) {
	totalNumber := c.PostForm("totalNumber")
	currentNumber := c.PostForm("currentNumber")
	fileMd5 := c.PostForm("fileMd5")
	fileHeader, err := c.FormFile("file")
	if err != nil {
		log.Panic(err)
	}

	tmpDir := dst + "tmp/" + fileMd5
	if currentNumber == "1" {
		os.MkdirAll(tmpDir, 0777)
	}

	// ch := make(chan string)

	// go func(cn chan string) {
	// 	cn <- "string"
	// }(ch)

	// <-chan

	// 保存当前分片
	c.SaveUploadedFile(fileHeader, tmpDir+"/"+fileMd5+"_"+currentNumber)

	// 最后一片
	if totalNumber == currentNumber {
		zipFile := tmpDir + "/" + fileHeader.Filename

		// 合并文件，校验md5值
		file, err := os.Open(tmpDir)
		if err != nil {
			log.Panicln(err)
		}
		fileInfos, _ := file.Readdir(0)
		for _, v := range fileInfos {
			tmpFile := tmpDir + "/" + v.Name()
			content, err := ioutil.ReadFile(tmpFile)
			if err != nil {
				panic(err)
			}
			ioutil.WriteFile(zipFile, content, 0755)
			os.Remove(tmpFile)

		}
		// 转移到目标目录
	}

}
