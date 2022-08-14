package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"not null" json:"name"`
	Password string `gorm:"not null" json:"password"`
	Type     int    `gorm:"not null" json:"type"`
}
