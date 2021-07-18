package config

import (
	"gin-boot/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", controller.UserList)

	return r
}
