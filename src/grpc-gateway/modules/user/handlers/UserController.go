package handlers

import (
	"fmt"
	"gin-gateway/conn"
	"gin-gateway/modules/user/models"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

// 查看
func (this *UserController) Index(c *gin.Context) {
	db := conn.GetDB()
	id := c.Query("id") // c.Request.URL.Query().Get("lastname") 的一种快捷方式'
	fmt.Println(id)
	// var user models.User
	var user []models.User
	// var results
	if id != "" {
		db.First(&user, id)
	} else {
		db.Find(&user)
	}
	c.JSON(200, user)
	c.JSON(200, map[string]string{
		"message": "dddddd",
	})
}

// 新增
func (this *UserController) Create(c *gin.Context) {
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
func (this *UserController) Update(c *gin.Context) {
	db := conn.GetDB()
	user := &models.User{}
	update_data := map[string]interface{}{
		"Name": "bbb",
		"Age":  30,
	}

	db.Model(user).Where("id = ?", "2").UpdateColumns(update_data)
	c.JSON(200, map[string]string{
		"message": "更新成功",
	})
}

// 删除
func (this *UserController) Delete(c *gin.Context) {
	db := conn.GetDB()
	id := c.Query("id")

	user := &models.User{}
	db.Where("id  = ?", id).Delete(user)
	c.JSON(200, map[string]string{
		"message": "删除成功",
	})
}
