package models

import (
	"log"
	"mygin/conn"
)

type BaseModel struct {
}

func (bm *BaseModel) FindOne(out interface{}, id string) interface{} {
	if err := conn.GetDb().First(out, id).Error; err != nil {
		log.Panic(err)
	}

	return out
}

func (bm *BaseModel) Add(data interface{}) {
	if err := conn.GetDb().Create(data).Error; err != nil {
		log.Panic(err)
	}
	// return data.id
}

func (bm *BaseModel) Edit(model interface{}, data interface{}) {
	if err := conn.GetDb().Model(model).Update(data).Error; err != nil {
		log.Panic(err)
	}
}

func(bm *BaseModel) Delete(value interface{}, where ...interface{}) {
	if err := conn.GetDb().Delete(value, where).Error; err != nil{
		log.Panic(err)
	}
}


func (bm *BaseModel) FindAll() {

}


