package db

import (
	"fmt"
	"log"
	"time"

	. "zjutjh/Join-Us/config"
	JHLog "zjutjh/Join-Us/utility/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	devLog := logger.New(
		log.New(JHLog.LogMultiWriter, "\r\n[GORM] ", log.LstdFlags),
		logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			Colorful:                  false,
			IgnoreRecordNotFoundError: false,
			LogLevel:                  logger.Info,
		})
	proLog := logger.New(
		log.New(JHLog.LogMultiWriter, "\r\n[GORM] ", log.LstdFlags),
		logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			Colorful:                  false,
			IgnoreRecordNotFoundError: false,
			LogLevel:                  logger.Warn,
		})
	username := Config.Database.Username
	password := Config.Database.Password
	address := Config.Database.Address
	dbname := Config.Database.DbName
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, address, dbname)
	var err error

	if Config.Dev {
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: devLog,
		})
	} else {
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: proLog,
		})
	}
	if err != nil {
		log.Panicln("Database Connection Error!", err)
	}

}
