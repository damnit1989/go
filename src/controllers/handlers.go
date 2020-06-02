package controllers

import (
	"mygin/utils"
	"time"

	"github.com/gin-gonic/gin"
)

var tokenStr = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTExMDY2OTksImp0aSI6InRlc3QiLCJpc3MiOiJsbW0ifQ.g0lTUU_bkDrzTpTmRgWSRmnxfCDCMI0sHS332QXr_z4"

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
