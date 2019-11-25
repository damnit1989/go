package routers

import (
	"beego_project/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/info", &controllers.MainController{})
	beego.Router("/test", &controllers.TestController{})
	beego.Router("/sql", &controllers.SqlController{})
}

