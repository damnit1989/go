package conn

import (
    "fmt"
    "time"
    "github.com/jinzhu/gorm"
    "gin-project/configs"
	// "gin-project/modules/user/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// 初始化sql连接
func init() {
	conn_str := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True&loc=Local", configs.DB_USER, configs.DB_PSWD, configs.DB_NAME, configs.DB_CHARSET)
    db, err := gorm.Open("mysql",conn_str)
    if err != nil{
        panic(err)
    }
    
    // 开启日志
    db.LogMode(true)

    // SetMaxIdleCons 设置连接池中的最大闲置连接数。
    db.DB().SetMaxIdleConns(configs.DB_MAX_IDLE_CONNS)

    // SetMaxOpenCons 设置数据库的最大连接数量。
    db.DB().SetMaxOpenConns(configs.DB_MAX_OPEN_CONNS)

    // SetConnMaxLifetiment 设置连接的最大可复用时间。
    db.DB().SetConnMaxLifetime(time.Hour)
}

func GetDB() *gorm.DB{
    return db
}