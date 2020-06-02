package models

// import "fmt"

type User struct {
	ID    int
	Name  string
	Age   int8
	Email string
	BaseModel
}

// func Info(id string) *User {
// 	var um User
// 	if err := conn.GetDb().Debug().First(&um, id).Error; err != nil {
// 		log.Panic(err)
// 	}
// 	// err := recover()

// 	return &um
// }

// func List() {

// }

// func Add(u User) {
// 	if err := conn.GetDb().Create(&u).Error; err != nil {
// 		log.Panic(err)
// 	}
// }

// func Update(u *User, data interface{}) {
// 	if err := conn.GetDb().Model(u).Where("status=?", 0).Update(data).Error; err != nil {
// 		log.Panic(err)
// 	}
// }

// func Delete(id string) {
// 	var um User
// 	if err := conn.GetDb().Delete(&um, id).Error; err != nil {
// 		log.Panic(err)
// 	}
// }
