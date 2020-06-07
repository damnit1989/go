package controllers

import (
	"log"
	"mygin/conn"
	"mygin/utils"
	"time"

	"github.com/gin-gonic/gin"
)

var tokenStr = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTEyMDA3NzQsImp0aSI6InRlc3QiLCJpc3MiOiJsbW0ifQ.wQ99MEy4-90FmzY-DJiKQgc5o357SDaP85ujtA25B1g"

func InfoPoint(c *gin.Context) {
	c.JSON(200, map[string]interface{}{
		"name": "lmm",
		"age":  23,
	})
}

func Login(c *gin.Context) {
	token := utils.GenerateToken()
	c.JSON(200, map[string]interface{}{
		"token": token,
		"msg":   "operate success",
	})
}

func VerifyToken(c *gin.Context) {
	token := utils.VerifyToken(tokenStr)
	c.JSON(200, gin.H{
		"token": token,
	})
}

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"times": time.Now().Unix(),
	})
}

type Queryparam struct {
	Name   string `form:"name" json:"name"  binding:"required"`
	Status int
}

func TestBindUrl(c *gin.Context) {
	var p Queryparam
	if err := c.ShouldBindQuery(&p); err != nil {
		log.Panic(err)
	}

	c.JSON(200, p)
}

func TestRedis(c *gin.Context) {
	cli := conn.GetRedis()
	cli.Set("name", "lmm", 0)
	val := cli.Get("name").Val()
	c.String(200, val)
	cli.Close()
}
