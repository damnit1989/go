package main

import (
	"fmt"
	"database/sql"
	"io"
	"log"
	"net/http"
	"text/template"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

//留言结构
type Liuyan struct{
	Id int
	Name string
	Content string
	Time int
}

//显示留言时间
func (l Liuyan) ShowTime() string{
	t := time.Unix(int64(l.Time), 0)
	return t.Format("2006-01-02 15:04:05")
}

func main(){
	db, err := sql.Open("mysql","root:123456@/test")
	if err != nil{
		panic(err)
	}
	defer db.Close()
	
	tpl, err := template.New("a.html").Parse("a.html")
	if err != nil{
		panic(err)
	}
	
	requestList := func(w http.ResponseWriter, req *http.Request){
		rows, err := db.Query("select * from user_log")
		if err != nil{
			log.Fatal(err)
		}
		defer rows.Close()
		
		lys := []Liuyan{}
		for rows.Next(){
			ly := Liuyan{}
			err := rows.Scan(&ly.Id, &ly.Name, &ly.Content, &ly.Time)
			if nil != err{
				log.Fatal(err)
			}
			lys = append(lys,ly)
		}
		
		err = tpl.ExecuteTemplate(w, "list", lys)
		if err != nil{
			log.Fatal(err)
		}
	}
	
	//留言页面
	requestLiuyan := func(w http.ResponseWriter, req *http.Request){
		err := req.ParseForm()
		if err != nil{
			log.Fatal(err)
		}
		
		if "POST" == req.Method {
			if len(req.Form["name"]) < 1 {
				io.WriteString(w,"参数错误\n")
				return 
			}
			if len(req.Form["content"]) < 1{
				io.WriteString(w, "参数错误\n")
				return 
			}
			name := template.HTMLEscapeString(req.Form.Get("name"))
			content := template.HTMLEscapeString(req.Form.Get("content"))
			
			//sql语句
			sql, err := db.Prepare("insert into liuyan(name,content,time) values(?, ?, ?)")
			if err != nil{
				log.Fatal(err)
			}
			defer sql.Close()
			
			res, err := sql.Exec(name,content,time.Now().Unix())
			if err != nil{
				log.Fatal(err)
			}
			fmt.Println(res)
			w.Header().Add("Location", "/")
			w.WriteHeader(302)
			io.WriteString(w, "提交成功\n")
			return
		}
		err = tpl.ExecuteTemplate(w, "liuyan", nil)
		if err != nil{
			log.Fatal(err)
		}
	}
	
	http.HandleFunc("/", requestList)
	http.HandleFunc("/liuyan", requestLiuyan)
	err = http.ListenAndServe("192.168.17.134:12345", nil)
	if err != nil{
		log.Fatal("ListenAndServe:", err)
	}
}