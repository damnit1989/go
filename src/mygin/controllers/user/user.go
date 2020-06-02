package user

import (
	"log"
	"mygin/conn"
	"mygin/models"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (u *UserController) InfoAction(c *gin.Context) {
	// id := c.Param("id")
	// // um := models.User{}
	// // id := 1
	// ret := models.Info(id)
	// c.JSON(200, ret)
}

func (u *UserController) DetailAction(c *gin.Context) {
	var um models.User
	id := c.Param("id")
	ret := um.FindOne(&um, id)
	c.JSON(200, ret)
}

func (u *UserController) AddAction(c *gin.Context) {
	name := c.PostForm("name")
	// age := c.PostForm("age")
	email := c.PostForm("email")
	um := models.User{
		Age:   40,
		Name:  name,
		Email: email,
	}
	um.Add(&um)

	c.JSON(200, map[string]interface{}{
		"message": "add success",
		"id":      um.ID,
	})
}

func (u *UserController) EditAction(c *gin.Context) {
	var um models.User
	id := c.PostForm("id")
	name := c.PostForm("name")
	// age := c.PostForm("age")
	age := 40
	email := c.PostForm("email")
	ret := um.FindOne(&um, id)
	data := map[string]interface{}{
		"name":  name,
		"age":   age,
		"email": email,
	}
	// models.Update(ret, data)
	um.Edit(ret, data)
	c.JSON(200, gin.H{
		"message": "updated sucdess",
	})
}

func (u *UserController) DeleteAction(c *gin.Context) {
	var um models.User
	id := c.Param("id")
	um.Delete(models.User{}, id)
	c.JSON(200, gin.H{
		"message": "delted sucdess",
	})
}

func (u *UserController) GetUserProfileAction(c *gin.Context) {
	id := c.Param("id")
	var um models.User
	up := new(models.UserProfile)
	ret := um.FindOne(&um, id)
	if err := conn.GetDb().Model(ret).Related(up).Error; err != nil {
		log.Panic(err)
	}
	// up.User = ret.(models.User)
	c.JSON(200, up)
}
