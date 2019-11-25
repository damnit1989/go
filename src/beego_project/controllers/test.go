package controllers

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
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
	c.ServeJSON()
}

func (c *TestController) Post() {
	id := c.GetString("id")
	name := c.GetString("name")
	if id == "" {
		data := map[string]interface{}{
			"code": 200,
			"msg":  "",
			"data": []string{},
		}
		c.Data["json"] = data
		c.ServeJSON()
		return
	}

	c.Ctx.WriteString(id)
	c.Ctx.WriteString(name)
	// c.Ctx.ResponseWriter.Write([]byte(id))
}
