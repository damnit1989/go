package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"beego_project/models"
)

type SqlController struct {
	beego.Controller
}

func (c *SqlController) Get() {

	o := orm.NewOrm()
	// o.Using("beego") // 默认使用 default，你可以指定为其他数据库

	user := new(models.User)
	user.Username = "slene"
	u, err := o.Insert(user)
	if err != nil {
		c.Ctx.WriteString("添加失败")
		return
	}

	file := new(models.File)
	file.Name = "临时文件"
	f, err := o.Insert(file)
	if err != nil {
		c.Ctx.WriteString("添加失败")
		return
	}

	data := map[string]interface{}{
		"code": 200,
		"msg":  "添加成功",
		"data": []int64{u, f},
	}
	c.Data["json"] = data
	c.ServeJSON()
}
