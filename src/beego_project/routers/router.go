package routers

import (
	"beego_project/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/info", &controllers.MainController{})
	// beego.Router("/test", &controllers.TestController{})
	beego.Router("/sql", &controllers.SqlController{})

	
	// 自动匹配路由：
	// beego 就会通过反射获取该结构体中所有的实现方法，你就可以通过如下的方式访问到对应的方法中：
	// /object/login   调用 ObjectController 中的 Login 方法
	// /object/logout  调用 ObjectController 中的 Logout 方法
	beego.AutoRouter(&controllers.TestController{})
}

