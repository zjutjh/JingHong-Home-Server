package model

import (
	"time"
	. "zjutjh/Join-Us/db"
)

type NormalForm struct {
	Name      string    `gorm:"not null" json:"name" xlsx:"姓名"`
	UpdatedAt time.Time `gorm:"not null" json:"-" xlsx:"更新时间"`
	StuID     string    `gorm:"primaryKey" json:"stu_id" xlsx:"学号"`
	Gender    string    `gorm:"not null" json:"gender" xlsx:"性别"`
	College   string    `gorm:"not null" json:"college" xlsx:"专业"`
	Campus    string    `gorm:"not null" json:"campus" xlsx:"学院"`
	Phone     string    `gorm:"not null" json:"phone" xlsx:"电话号"`
	QQ        string    `gorm:"not null" json:"qq" xlsx:"QQ"`
	Region    int       `gorm:"not null" json:"region" xlsx:"校区"`
	Want1     int       `gorm:"not null" json:"want1" xlsx:"第一志愿"`
	Want2     int       `gorm:"not null" json:"want2" xlsx:"第二志愿"`
	Profile   string    `gorm:"not null" json:"profile" xlsx:"简介"`
	Feedback  string    `gorm:"not null" json:"feedback" xlsx:"反馈"`
}

func NewNormalForm(form *NormalForm) error {
	return DB.Create(form).Error
}

func UpdateNormalForm(form *NormalForm) error {
	return DB.Save(form).Error
}

func GetNormalFormByStuID(stuID string) (*NormalForm, error) {
	var form NormalForm
	form.StuID = stuID
	err := DB.First(&form).Error
	return &form, err
}

func GetAllNormalForms() ([]NormalForm, error) {
	forms := []NormalForm{}
	err := DB.Find(&forms).Error
	return forms, err
}

type FormTotal [][]int

const (
	want1    = 0
	want1ZH  = 1
	want1PF  = 2
	want1MGS = 3
	want2    = 4
	want2ZH  = 5
	want2PF  = 6
	want3MGS = 7
	Total    = 8
	Today    = 9
)

func GetNormalFormTotal() FormTotal {
	forms, _ := GetAllNormalForms()
	res := FormTotal{{0, 0, 0, 0, 0}}
	for i := 0; i < 9; i += 1 {
		res = append(res, make([]int, 10))
	}
	today := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)
	for _, v := range forms {
		// 总计赋值
		res[0][0] += 1
		if v.UpdatedAt.Unix() >= today.Unix() {
			res[0][1] += 1
			res[v.Want1][Today] += 1
			if v.Want2 != 0 {
				res[v.Want2][Today] += 1
			}
		}
		res[0][v.Region+1] += 1
		// 分部门统计
		res[v.Want1][want1] += 1
		res[v.Want1][v.Region] += 1
		res[v.Want1][Total] += 1
		if v.Want2 != 0 {
			res[v.Want2][want2] += 1
			res[v.Want2][v.Region+4] += 1
			res[v.Want2][Total] += 1
		}
	}
	return res
}

type NormalFormBrief struct {
	Name      string    `gorm:"not null" json:"name"`
	StuID     string    `gorm:"primaryKey" json:"stu_id"`
	Want1     int       `gorm:"not null" json:"want1"`
	Want2     int       `gorm:"not null" json:"want2"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}

func GetAllNormalFormsBrief() ([]NormalFormBrief, error) {
	forms := []NormalFormBrief{}
	err := DB.Model(&NormalForm{}).Find(&forms).Error
	return forms, err
}
