package core

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GORM *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.dsn"),
	)

	dbConfig := gorm.Config{}

	var err error
	GORM, err = gorm.Open(mysql.Open(dsn), &dbConfig)
	if err != nil {
		log.Fatal("InitDB fail:" + err.Error())
	}

	db, err := GORM.DB()
	db.SetMaxIdleConns(viper.GetInt("mysql.maxIdleConns"))
	db.SetMaxOpenConns(viper.GetInt("mysql.maxOpenConns"))

	return
}
