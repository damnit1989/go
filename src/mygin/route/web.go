package route

import (
	"mygin/controllers"
	"mygin/controllers/upload"
	"mygin/controllers/user"
	"mygin/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var accounts = gin.Accounts{
	"lmm": "123456",
	"abc": "123456",
}

func init() {
	router = gin.Default()

}

// SetRoute return route
func SetRoute() *gin.Engine {

	// 简单的路由组: v1
	v1 := router.Group("/v1", middlewares.AuthToken())
	{
		u := user.UserController{}
		v1.GET("/index/:id", (&u).InfoAction)
		v1.GET("/detail/:id", (&u).DetailAction)
		v1.POST("/add", (&u).AddAction)
		v1.POST("/edit", (&u).EditAction)
		v1.GET("/delete/:id", (&u).DeleteAction)
		v1.GET("/get-user-profile/:id", (&u).GetUserProfileAction)
	}

	// 简单的路由组: v2
	v2 := router.Group("/v2")
	{
		v2.GET("/a", controllers.InfoPoint)
	}

	// 上传文件
	file := router.Group("/upload")
	{
		file.POST("/pic", upload.UploadPic)
		file.POST("/breakpoint-upload", upload.BreakPointUpload)
	}

	// 后端
	admin := router.Group("/admin", middlewares.Test(), gin.BasicAuth(accounts))
	{
		admin.GET("/user", func(c *gin.Context) {
			age := c.MustGet("age").(int)
			c.JSON(http.StatusOK, gin.H{"age": age})
			user := c.MustGet(gin.AuthUserKey).(string)
			if secret, ok := accounts[user]; ok {
				c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
			} else {
				c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
			}
		})
	}

	router.GET("/login/login", controllers.Login)
	router.GET("/verify-token", controllers.VerifyToken)
	router.GET("/test", controllers.Test)
	router.GET("/bind-url", controllers.TestBindUrl)
	router.GET("/redis", controllers.TestRedis)

	return router
}
