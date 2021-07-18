package service

import "gin-boot/model"

func UserList() (users []model.User, err error) {
	user := model.User{}
	return user.Users()
}
