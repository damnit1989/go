package main

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"gin-project/conn"
	"github.com/jinzhu/gorm"
	"gin-project/modules/user/handlers"
	"gin-project/modules/user/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "database/sql"
	// "time"
)


func create(db *gorm.DB) {
	user := models.User{
		Name:  "lmm",
		Email: "448559829@qq.com",
		Age:   20,
	}
	db.Create(&user)
}

func update(db *gorm.DB, user models.User) {

}

func findByWhereMap(db *gorm.DB, where map[string]interface{},user models.User) {
	
}

func findByWhereStruct(db *gorm.DB,where struct{},user models.User){

}

func delete(db *gorm.DB, user models.User) {
	db.Delete(&user)
}

func main() {
	db := conn.GetDB()

	defer db.Close()

	//设置默认表名前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "gin_" + defaultTableName
	}

	// fmt.Println(db.GetErrors())
	var u handlers.UserController
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context){
		c.JSON(401,gin.H{"message":"pong",
		})
	})
	router.GET("/")
	router.GET("/index", u.Index)
	router.GET("/update", u.Update)
	router.GET("/create", u.Create)
	router.GET("/delete", u.Delete)
	router.Run("0.0.0.0:8312")
}
