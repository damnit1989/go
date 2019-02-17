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
	rows, err := db.Query("SELECT name FROM user_log WHERE id=?", id)
	if err != nil{
		log.Fatal(err)
	}
	
	defer rows.Close()
	
	for rows.Next(){
		var name string
		if err := rows.Scan(&name);err != nil{
			log.Fatal(err)
		}
		fmt.Printf("%s is %d\n",name,id)
	}
	if err := rows.Err();err != nil{
		log.Fatal(err)
	}

}