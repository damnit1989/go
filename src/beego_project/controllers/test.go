package controllers

import (
	"beego_project/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
	// params := c.Ctx.Input.Params()
	id := c.Ctx.Input.Query("id")
	name := c.Ctx.Input.Query("name")

	// panic(params)
	data := map[string]interface{}{
		"code": 200,
		"msg":  "",
		"data": []map[string]interface{}{
			{"name": "aa", "age": 23},
			{"name": "aa", "age": 23},
			{"name": "aa", "age": 23},
		},
	}
	c.Data["json"] = data
	c.Data["json"] = map[string]string{
		"id":   id,
		"name": name,
	}
	c.ServeJSON()
}

func (c *TestController) Post() {
	id := c.GetString("id")
	name := c.GetString("name")
	if id != "" {
		data := map[string]interface{}{
			"code": 200,
			"msg":  "获取成功",
			"data": map[string]string{
				"id":   id,
				"name": name,
			},
		}
		c.Data["json"] = data
		c.ServeJSON()
		return
	}

	c.Ctx.WriteString(id)
	c.Ctx.WriteString(name)
	// c.Ctx.ResponseWriter.Write([]byte(id))
}

func (c *TestController) Info() {
	id := c.Ctx.Input.Query("id")
	if id == "" {
		panic("参数错误")
	}
	o := orm.NewOrm()
	user := models.User{Id: 1}

	err := o.Read(&user)

	if err == orm.ErrNoRows {
		panic("查询不到")
	} else if err == orm.ErrMissPK {
		panic("找不到主键")
	} else {
		c.Data["json"] = user
		c.ServeJSON()
		c.Ctx.WriteString(user.Username)
	}
	return
}

func (c *TestController) List() {

	if bool := c.Ctx.Input.IsAjax(); bool == false {
		c.Data["json"] = map[string]interface{}{
			"code": 403,
			"msg":  "非法请求，请使用ajax请求",
		}
		c.ServeJSON()
	}

	o := orm.NewOrm()
	var users []*models.User
	// 也可以直接使用对象作为表名
	// user := new(models.User)
	_, err := o.QueryTable("bg_user").Filter("id__gte", 1).All(&users) // WHERE id >= 1
	if err != nil {
		panic("操作失败")
	}
	c.Data["json"] = users
	c.ServeJSON()
}

func(c *TestController) Delete(){
	defer func(){
		if r := recover(); r != nil{
			c.Data["json"] = map[string]interface{}{
				"code": 403,
				"msg":  r,
			}
			c.ServeJSON()			
		}
	}()

	o := orm.NewOrm()
	if bool := c.Ctx.Input.IsAjax(); bool == false {
		panic("非法请求，请使用ajax请求")
	}

	if id,err:= c.GetInt("id");err != nil{
		panic("参数错误")
	}

	if _, err := o.Delete(&models.User{Id: id}); err != nil {
		panic("删除失败")
	}
	
	c.Data["json"] = map[string]interface{}{
		"code": 200,
		"msg":  "删除成功",
	}
	c.ServeJSON()	
}
