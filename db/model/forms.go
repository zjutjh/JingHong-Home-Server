package model

import "time"

type NormalForm struct {
	Name      string    `gorm:"not null" json:"name"`
	UpdatedAt time.Time `gorm:"not null" json:"updated"`
	StuID     string    `gorm:"primaryKey" json:"stu_id"`
	Gender    string    `gorm:"not null" json:"gender"`
	College   string    `gorm:"not null" json:"college"`
	Campus    string    `gorm:"not null" json:"campus"`
	Phone     string    `gorm:"not null" json:"phone"`
	QQ        string    `gorm:"not null" json:"qq"`
	Region    string    `gorm:"not null" json:"region"`
	Want1     string    `gorm:"not null" json:"want1"`
	Want2     string    `gorm:"not null" json:"want2"`
	Profile   string    `gorm:"not null" json:"profile"`
	Feedback  string    `gorm:"not null" json:"feedback"`
}
