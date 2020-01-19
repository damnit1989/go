package conn

import (
	"fmt"
	"gin-project/configs"
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// 初始化数据库连接
func init() {
	var err error
	connStr := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True&loc=Local", configs.DB_USER, configs.DB_PSWD, configs.DB_NAME, configs.DB_CHARSET)
	db, err = gorm.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}

	// 添加表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "gin_" + defaultTableName
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

// GetDB 获取数据库连接
func GetDB() *gorm.DB {
	return db
}
