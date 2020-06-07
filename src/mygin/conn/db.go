package conn

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB_USER    = "root"
	DB_PSWD    = "12345678"
	DB_NAME    = "my_gin"
	DB_CHARSET = "utf8"
)

var db *gorm.DB

func init() {
	connStr := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True&loc=Local", DB_USER, DB_PSWD, DB_NAME, DB_CHARSET)
	conn, err := gorm.Open("mysql", connStr)
	if err != nil {
		// log.Fatal("db conn failed")
		log.Panic("db conn failed")
	}

	db = conn

	// 开启db调试，可打印执行sql
	db.LogMode(true)

	// defer db.Close()
}

// GetDb return db conn
func GetDb() *gorm.DB {
	return db
}
