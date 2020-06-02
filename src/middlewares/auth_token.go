package middlewares

import (
	"mygin/utils"

	"github.com/gin-gonic/gin"
)

func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		// c.String(200, "%s", token)
		utils.VerifyToken(token)
		c.Abort()
	}
}
