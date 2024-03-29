package model

type DeveloperForm struct {
	Name         string  `gorm:"not null" json:"name"`
	StuID        string  `gorm:"primaryKey" json:"stu_id"`
	Gender       string  `gorm:"not null" json:"gender"`
	College      string  `gorm:"not null" json:"college"`
	Campus       string  `gorm:"not null" json:"campus"`
	Phone        string  `gorm:"not null" json:"phone"`
	QQ           string  `gorm:"not null" json:"qq"`
	Region       string  `gorm:"not null" json:"region"`
	Profile      string  `gorm:"not null" json:"profile"`
	Feedback     string  `gorm:"not null" json:"feedback"`
	AbilityOther string  `gorm:"not null" json:"ability_other"`
	Ability      Ability `json:"ability"`
}

type Ability struct {
	DeveloperFormID uint `json:"-"`
	Api             bool `json:"api"`
	FrontEnd        bool `json:"front_end"`
	Document        bool `json:"document"`
	Git             bool `json:"git"`
}
