package utils

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"sort"
)

type FileInfo struct {
	Size int64
	Path string
	Md5  string
}

// Upload is the struct
type Upload struct {
	FileType   []string
	FileSzie   int64
	UploadFold string
}

// CheckMimeType check the upload file mime-type
func (u *Upload) CheckMimeType(file *multipart.FileHeader) {
	contentType := file.Header.Get("Content-Type")
	sort.Strings(u.FileType)
	pos := sort.SearchStrings(u.FileType, contentType)
	if pos == len(u.FileType) {
		log.Panic("文件mime-type不正确")
	}
	if contentType != u.FileType[pos] {
		log.Panic("文件mime-type不正确")
	}
}

// CheckFileSize check the upload file'size
func (u *Upload) CheckFileSize(file *multipart.FileHeader) {
	size := file.Size
	if size > u.FileSzie {
		log.Panic("文件太大，请上传10M一下的文件")
	}
}

// CheckFoldExist check the target fold if exist or not
func (u *Upload) CheckFoldExist() {
	fileInfo, err := os.Stat(u.UploadFold)
	if err != nil {
		log.Panic("目录检查错误")
	}
	if !fileInfo.IsDir() {
		log.Panic("目录不存在")
	}
}

// DoUpload do upload file
func (u *Upload) DoUpload(file *multipart.FileHeader) *FileInfo {
	u.CheckMimeType(file)
	u.CheckFileSize(file)
	u.CheckFoldExist()

	src, err := file.Open()
	if err != nil {
		log.Panic("文件读取失败")
	}

	defer src.Close()

	filePath := u.UploadFold + file.Filename
	dst, err := os.Create(filePath)
	if err != nil {
		log.Panic("目标文件创建失败")
	}

	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		log.Panic("上传失败")
	}

	return &FileInfo{
		Size: file.Size,
		Path: filePath,
		Md5:  FileMd5(filePath),
	}
}

// NewUpload return upload instance
func NewUpload(config map[string]interface{}) *Upload {
	uploadIns := &Upload{
		FileType:   []string{"image/jpeg", "image/png"},
		FileSzie:   1024 * 1024 * 10,
		UploadFold: "./uploads/",
	}
	if fileType, ok := config["FileType"]; ok {
		uploadIns.FileType = fileType.([]string)
	}
	if fileSize, ok := config["FileSize"]; ok {
		uploadIns.FileSzie = fileSize.(int64)
	}
	if uploadFold, ok := config["UploadFold"]; ok {
		uploadIns.UploadFold = uploadFold.(string)
	}
	return uploadIns
}
