//测试数据库连接
package main

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
)

func main(){
	db,err := sql.Open("mysql","root:123456@/test")
	if err != nil{
		//fmt.Printf("ddd");
		log.Fatal(err)

	}
	id := 363
	rows, err := db.Query("SELECT name,content  FROM user_log WHERE id >= ?", id)
	if err != nil{
		log.Fatal(err)
	}
	
	defer rows.Close()
	
	for rows.Next(){
		var name, content string
		if err := rows.Scan(&name, &content);err != nil{
			log.Fatal(err)
		}
		fmt.Printf("name:%s,content:%s,id:%d\n",name, content, id)
	}
	if err := rows.Err();err != nil{
		log.Fatal(err)
	}

}
