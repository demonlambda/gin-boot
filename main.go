package main

import (
	"fmt"
	"gin-boot/config"
	"gin-boot/core"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

//全局变量，支持配置文件设置
var (
	port string
)

var (
	db *gorm.DB
)

func main() {
	core.InitDB()
	db, _ := core.GORM.DB()
	defer db.Close()

	r := config.InitRouter()
	r.Run(":" + port)
}

//根据启动参数中指定的环境信息，确定启动的端口
func init() {
	port = viper.GetString("port")
	fmt.Printf("bootstrap config port:%s\n", port)
}
