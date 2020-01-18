package controllers

import (
	"github.com/astaxie/beego"
	// "fmt"
)

// type MainController struct {
// 	beego.Controller
// }

type TestController struct{
	beego.Controller
}

type Content struct{
	Data []map[string]interface{}
	Msg string
	Code int 
}


func (c *TestController) Get() {
	// mystruct := Content{Name : "lmm",Age : 23}
	mystruct := Content{
		Data:[]map[string]interface{}{
			{"id":12,"name":"aa"},
			{"id":13,"name":"bb"},
		},
		Msg:"success",
		Code:0,
	}
	c.Data["json"] = &mystruct
	c.ServeJSON()
}

func(c *TestController) Post(){

}