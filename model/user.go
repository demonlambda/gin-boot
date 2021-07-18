package model

import "gin-boot/core"

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func (User) TableName() string {
	return "user"
}

func (user *User) Users() (users []User, err error) {
	if err = core.GORM.Find(&users).Error; err != nil {
		return
	}
	return
}
