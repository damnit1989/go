package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
    // 需要在init中注册定义的model
	user := beego.AppConfig.String("mysqluser")
	pass := beego.AppConfig.String("mysqlpass")
	// urls := beego.AppConfig.String("mysqlurls")
	db := beego.AppConfig.String("mysqldb")
	prefix := beego.AppConfig.String("mysqlprefix")

	// orm.RegisterDataBase("default", "mysql", "root:123456@/beego?charset=utf8")
	orm.RegisterDataBase("default", "mysql", user+":"+pass+"@/"+db+"?charset=utf8")
	orm.RegisterModelWithPrefix(prefix, new(User),new(File))
	orm.Debug = true
}