package model

import (
	"log"
	. "zjutjh/Join-Us/db"
)

func init() {
	err := DB.AutoMigrate(NormalForm{})
	if err != nil {
		log.Panicln("Database Create Lists Error!", err)
	}
}
