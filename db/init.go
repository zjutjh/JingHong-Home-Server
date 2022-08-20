package db

import (
	"fmt"
	"log"

	. "zjutjh/Join-Us/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	username := Config.Database.Username
	password := Config.Database.Password
	address := Config.Database.Address
	dbname := Config.Database.DbName
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, address, dbname)
	var err error
	if Config.Dev {
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	} else {
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}
	if err != nil {
		log.Panicln("Database Connection Error!", err)
	}

}
