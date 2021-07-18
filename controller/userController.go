package controller

import (
	"gin-boot/service"

	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	userList, err := service.UserList()
	if err == nil {
		c.JSON(200, gin.H{"error": 0, "msg": "", "data": userList})
	} else {
		c.JSON(200, gin.H{"error": 0, "msg": "服务器繁忙", "data": nil})
	}

	return
}
