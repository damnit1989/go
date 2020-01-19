package main

import (
	// "fmt"
	"gin-project/modules/user/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "time"
)

func main() {
	// db := conn.GetDB()
	// defer db.Close()

	//设置默认表名前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "gin_" + defaultTableName
	}

	// fmt.Println(db.GetErrors())
	var u handlers.UserController
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(401, gin.H{"message": "pong"})
	})
	router.GET("/")
	router.GET("/index", u.Index)
	router.GET("/update", u.Update)
	router.GET("/create", u.Create)
	router.GET("/delete", u.Delete)
	router.Run("0.0.0.0:8312")
}
