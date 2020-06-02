package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Test() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("age", 23)
	}
}
