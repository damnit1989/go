package handlers

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"gin-project/conn"
	"gin-project/modules/user/models"
)

type UserController struct{}

// 查看
func(this *UserController) Index(c *gin.Context) {
	// db := conn.GetDB()
	// id := c.Query("id") // c.Request.URL.Query().Get("lastname") 的一种快捷方式'
	// var user models.User
	// // var results 
	// if id != "" {
	// 	db.First(&user, id).Row()
	// } else {
	// 	db.Find(&user)
	// }
	// c.JSON(200, user)
	c.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API",
	})
}

// 新增
func(this *UserController) Create(c *gin.Context) {
    db := conn.GetDB()
    // panic(db)
    user := models.User{
		Name:  "lmm",
		Email: "448559829@qq.com",
		Age:   20,
	}
    db.Create(&user)
    
	c.JSON(200, map[string]string{
		"message": "添加成功",
	})
}

// 更新
func(this *UserController) Update(c *gin.Context) {
	c.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API",
	})
}

// 删除
func(this *UserController) Delete(c *gin.Context) {
	c.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API",
	})
}
