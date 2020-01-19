package handlers

import (
<<<<<<< HEAD
	"fmt"
=======
	// "fmt"
>>>>>>> b69b84f93ef66102c83a35a6ac02a7d6e5665b8d
	"github.com/gin-gonic/gin"
	"gin-project/conn"
	"gin-project/modules/user/models"
)

type UserController struct{}

// 查看
func(this *UserController) Index(c *gin.Context) {
<<<<<<< HEAD
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
=======
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
>>>>>>> b69b84f93ef66102c83a35a6ac02a7d6e5665b8d
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
<<<<<<< HEAD
	db := conn.GetDB()
	user := &models.User{}
	update_data := map[string]interface{}{
		"Name": "bbb",
		"Age" : 30,
	}

	db.Model(user).Where("id = ?", "2").UpdateColumns(update_data)
	c.JSON(200, map[string]string{
		"message": "更新成功",
=======
	c.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API",
>>>>>>> b69b84f93ef66102c83a35a6ac02a7d6e5665b8d
	})
}

// 删除
func(this *UserController) Delete(c *gin.Context) {
<<<<<<< HEAD
	db := conn.GetDB()
	id := c.Query("id")

	user := &models.User{}
	db.Where("id  = ?", id).Delete(user)
	c.JSON(200, map[string]string{
		"message": "删除成功",
=======
	c.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API",
>>>>>>> b69b84f93ef66102c83a35a6ac02a7d6e5665b8d
	})
}
