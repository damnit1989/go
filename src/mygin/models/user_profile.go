package models

// import "fmt"

type UserProfile struct {
	ID     int
	Name   string
	User   User `gorm:"association_foreignkey:UserID"` // 将 UserRefer 作为外键
	UserID int
	BaseModel
}

func (UserProfile) TableName() string {
	return "user_profiles"
}
