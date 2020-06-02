package main

import (
	"io"
	"mygin/route"
	"os"

	"github.com/gin-gonic/gin"
)

// import "fmt"

func main() {
	// r := gin.Default()
	f, _ := os.Create("runtime/info.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := route.SetRoute()
	// gin.BasicAuth()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
