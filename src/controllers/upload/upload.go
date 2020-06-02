package upload

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// const UPLOAD_PATH = "../../uploads/"

var dst = "./uploads/"

func UploadPic(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Panicln(err)
	}
	// c.JSON(200, file.Header)
	// fileInfo, err := file.Open()

	c.SaveUploadedFile(file, dst+file.Filename)
	c.String(200, "file %s  upload success", file.Filename)
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

			// pieceFile, _ := os.Open(tmpDir + "/" + v.Name())
			// pieceFile.Read()
			// pieceFile.Read(pieceFile.
		}
		// 转移到目标目录
	}

}
