package model

type User struct {
	UserName string `gorm:"not null" json:"name"`
	Password string `gorm:"not null" json:"password"`
	Type     int    `gorm:"not null" json:"type"`
}
