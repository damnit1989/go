package models

type User struct{
	Id int
	Username string
	Nick_name string
	Email string
}

func init() {
    // 需要在init中注册定义的model
}
