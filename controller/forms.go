package controller

import (
	"fmt"
	"time"
	"zjutjh/Join-Us/model"
	"zjutjh/Join-Us/utility"
	"zjutjh/Join-Us/utility/initial"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func NewNormalForm(c *gin.Context) {
	postData := model.NormalForm{}
	err := c.ShouldBindBodyWith(&postData, binding.JSON)
	if err != nil {
		fmt.Println(err)
		utility.ResponseError(c, "Post Data Error")
	}
	res := initial.DB.Save(&postData)
	if res.RowsAffected == 0 {
		utility.ResponseError(c, "Database Error")
	} else {
		utility.ResponseSuccess(c, nil)
	}
}

func NewDevelopForm(c *gin.Context) {
	data := model.DeveloperForm{}
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utility.ResponseError(c, "Json error")
	}
	raw_data := model.DeveloperForm{}
	raw_data.StuID = data.StuID
	initial.DB.Where("stu_id = ?", data.StuID).First(&raw_data)
	if (raw_data == model.DeveloperForm{}) {
		initial.DB.Save(&data)
	} else {
		aff := initial.DB.Create(&data)

		if aff.RowsAffected == 0 {
			utility.ResponseError(c, "Had been Posted")
		} else {
			utility.SendEmail(initial.Config.Email.Receiver)
			utility.ResponseSuccess(c, nil)
		}
	}
}

func GetAllNormalForms(c *gin.Context) {
	forms := []model.NormalForm{}
	initial.DB.Find(&forms)
	utility.ResponseSuccess(c, gin.H{
		"forms": forms,
	})
}

func GetAllDevelopForms(c *gin.Context) {
	dev := []model.DeveloperForm{}
	initial.DB.Preload("Ability").Find(&dev)
	utility.ResponseSuccess(c, gin.H{
		"forms": dev,
	})
}

type departmentsData struct {
	Total    int `json:"total"`
	Today    int `json:"today"`
	Want1    int `json:"want1"`
	Want2    int `json:"want2"`
	Want1Zh  int `json:"want1_zh"`
	Want2Zh  int `json:"want2_zh"`
	Want1Pf  int `json:"want1_pf"`
	Want2Pf  int `json:"want2_pf"`
	Want1Mgs int `json:"want1_mgs"`
	Want2Mgs int `json:"want2_mgs"`
}
type FormsDataTotal struct {
	Total      int             `json:"total"`
	TotalToday int             `json:"total_today"`
	TotalZH    int             `json:"total_zh"`
	TotalPF    int             `json:"total_pf"`
	TotalMGS   int             `json:"total_mgs"`
	Bgs        departmentsData `json:"bgs"`
	Hdb        departmentsData `json:"hdb"`
	Msc        departmentsData `json:"msc"`
	Touch      departmentsData `json:"touch"`
	Xh         departmentsData `json:"xh"`
	Bj         departmentsData `json:"bj"`
	Sj         departmentsData `json:"sj"`
	Kfb        departmentsData `json:"kfb"`
	Yb         departmentsData `json:"yb"`
}

func GetFormsTotal(c *gin.Context) {
	forms := []model.NormalForm{}
	initial.DB.Find(&forms)
	f := FormsDataTotal{}
	f.Total = len(forms)
	// fmt.Println(forms)
	today := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)

	for i := 0; i < len(forms); i++ {
		var isToday bool = false
		if forms[i].UpdatedAt.Unix() >= today.Unix() {
			f.TotalToday += 1
			isToday = true
		}

		switch forms[i].Want1 {
		case "?????????":
			f.Bgs.Total += 1
			if isToday {
				f.Bgs.Today += 1
			}
			f.Bgs.Want1 += 1
			if forms[i].Region == "??????" {
				f.Bgs.Want1Zh += 1
			} else if forms[i].Region == "??????" {
				f.Bgs.Want1Pf += 1
			} else if forms[i].Region == "?????????" {
				f.Bgs.Want1Mgs += 1
			}
			break
		case "?????????":
			f.Hdb.Total += 1
			if isToday {
				f.Hdb.Today += 1
			}
			f.Hdb.Want1 += 1
			if forms[i].Region == "??????" {
				f.Hdb.Want1Zh += 1
			} else if forms[i].Region == "??????" {
				f.Hdb.Want1Pf += 1
			} else if forms[i].Region == "?????????" {
				f.Hdb.Want1Mgs += 1
			}
			break
		case "?????????":
			f.Msc.Total += 1
			if isToday {
				f.Msc.Today += 1
			}
			f.Msc.Want1 += 1
			if forms[i].Region == "??????" {
				f.Msc.Want1Zh += 1
			} else if forms[i].Region == "??????" {
				f.Msc.Want1Pf += 1
			} else if forms[i].Region == "?????????" {
				f.Msc.Want1Mgs += 1
			}
			break
		case "Touch?????????":
			f.Touch.Total += 1
			if isToday {
				f.Touch.Today += 1
			}
			f.Touch.Want1 += 1
			if forms[i].Region == "??????" {
				f.Touch.Want1Zh += 1
			} else if forms[i].Region == "??????" {
				f.Touch.Want1Pf += 1
			} else if forms[i].Region == "?????????" {
				f.Touch.Want1Mgs += 1
			}
			break
		case "???????????????":
			f.Xh.Total += 1
			if isToday {
				f.Xh.Today += 1
			}
			f.Xh.Want1 += 1
			if forms[i].Region == "??????" {
				f.Xh.Want1Zh += 1
			} else if forms[i].Region == "??????" {
				f.Xh.Want1Pf += 1
			} else if forms[i].Region == "?????????" {
				f.Xh.Want1Mgs += 1
			}
			break
		case "???????????????":
			f.Bj.Total += 1
			if isToday {
				f.Bj.Today += 1
			}
			f.Bj.Want1 += 1
			if forms[i].Region == "??????" {
				f.Bj.Want1Zh += 1
			} else if forms[i].Region == "??????" {
				f.Bj.Want1Pf += 1
			} else if forms[i].Region == "?????????" {
				f.Bj.Want1Mgs += 1
			}
			break
		case "???????????????":
			f.Sj.Total += 1
			if isToday {
				f.Sj.Today += 1
			}
			f.Sj.Want1 += 1
			if forms[i].Region == "??????" {
				f.Sj.Want1Zh += 1
			} else if forms[i].Region == "??????" {
				f.Sj.Want1Pf += 1
			} else if forms[i].Region == "?????????" {
				f.Sj.Want1Mgs += 1
			}
			break
		case "?????????":
			f.Kfb.Total += 1
			if isToday {
				f.Kfb.Today += 1
			}
			f.Kfb.Want1 += 1
			if forms[i].Region == "??????" {
				f.Kfb.Want1Zh += 1
			} else if forms[i].Region == "??????" {
				f.Kfb.Want1Pf += 1
			} else if forms[i].Region == "?????????" {
				f.Kfb.Want1Mgs += 1
			}
			break
		case "?????????????????????":
			f.Yb.Total += 1
			if isToday {
				f.Yb.Today += 1
			}
			f.Yb.Want1 += 1
			if forms[i].Region == "??????" {
				f.Yb.Want1Zh += 1
			} else if forms[i].Region == "??????" {
				f.Yb.Want1Pf += 1
			} else if forms[i].Region == "?????????" {
				f.Yb.Want1Mgs += 1
			}
			break
		}

		switch forms[i].Want2 {
		case "?????????":
			f.Bgs.Total += 1
			if isToday {
				f.Bgs.Today += 1
			}
			f.Bgs.Want2 += 1
			if forms[i].Region == "??????" {
				f.Bgs.Want2Zh += 1
			} else if forms[i].Region == "??????" {
				f.Bgs.Want2Pf += 1
			} else if forms[i].Region == "?????????" {
				f.Bgs.Want2Mgs += 1
			}
			break
		case "?????????":
			f.Hdb.Total += 1
			if isToday {
				f.Hdb.Today += 1
			}
			f.Hdb.Want2 += 1
			if forms[i].Region == "??????" {
				f.Hdb.Want2Zh += 1
			} else if forms[i].Region == "??????" {
				f.Hdb.Want2Pf += 1
			} else if forms[i].Region == "?????????" {
				f.Hdb.Want2Mgs += 1
			}
			break
		case "?????????":
			f.Msc.Total += 1
			if isToday {
				f.Msc.Today += 1
			}
			f.Msc.Want2 += 1
			if forms[i].Region == "??????" {
				f.Msc.Want2Zh += 1
			} else if forms[i].Region == "??????" {
				f.Msc.Want2Pf += 1
			} else if forms[i].Region == "?????????" {
				f.Msc.Want2Mgs += 1
			}
			break
		case "Touch?????????":
			f.Touch.Total += 1
			if isToday {
				f.Touch.Today += 1
			}
			f.Touch.Want2 += 1
			if forms[i].Region == "??????" {
				f.Touch.Want2Zh += 1
			} else if forms[i].Region == "??????" {
				f.Touch.Want2Pf += 1
			} else if forms[i].Region == "?????????" {
				f.Touch.Want2Mgs += 1
			}
			break
		case "???????????????":
			f.Xh.Total += 1
			if isToday {
				f.Xh.Today += 1
			}
			f.Xh.Want2 += 1
			if forms[i].Region == "??????" {
				f.Xh.Want2Zh += 1
			} else if forms[i].Region == "??????" {
				f.Xh.Want2Pf += 1
			} else if forms[i].Region == "?????????" {
				f.Xh.Want2Mgs += 1
			}
			break
		case "???????????????":
			f.Bj.Total += 1
			if isToday {
				f.Bj.Today += 1
			}
			f.Bj.Want2 += 1
			if forms[i].Region == "??????" {
				f.Bj.Want2Zh += 1
			} else if forms[i].Region == "??????" {
				f.Bj.Want2Pf += 1
			} else if forms[i].Region == "?????????" {
				f.Bj.Want2Mgs += 1
			}
			break
		case "???????????????":
			f.Sj.Total += 1
			if isToday {
				f.Sj.Today += 1
			}
			f.Sj.Want2 += 1
			if forms[i].Region == "??????" {
				f.Sj.Want2Zh += 1
			} else if forms[i].Region == "??????" {
				f.Sj.Want2Pf += 1
			} else if forms[i].Region == "?????????" {
				f.Sj.Want2Mgs += 1
			}
			break
		case "?????????":
			f.Kfb.Total += 1
			if isToday {
				f.Kfb.Today += 1
			}
			f.Kfb.Want2 += 1
			if forms[i].Region == "??????" {
				f.Kfb.Want2Zh += 1
			} else if forms[i].Region == "??????" {
				f.Kfb.Want2Pf += 1
			} else if forms[i].Region == "?????????" {
				f.Kfb.Want2Mgs += 1
			}
			break

		case "?????????????????????":
			f.Yb.Total += 1
			if isToday {
				f.Yb.Today += 1
			}
			f.Yb.Want2 += 1
			if forms[i].Region == "??????" {
				f.Yb.Want2Zh += 1
			} else if forms[i].Region == "??????" {
				f.Yb.Want2Pf += 1
			} else if forms[i].Region == "?????????" {
				f.Yb.Want2Mgs += 1
			}
			break
		}
	}

	f.Total = f.Bgs.Total + f.Hdb.Total + f.Msc.Total + f.Touch.Total + f.Xh.Total + f.Bj.Total + f.Sj.Total + f.Kfb.Total + f.Yb.Total
	f.TotalToday = f.Bgs.Today + f.Hdb.Today + f.Msc.Today + f.Touch.Today + f.Xh.Today + f.Bj.Today + f.Sj.Today + f.Kfb.Today + f.Yb.Today

	f.TotalZH = f.Bgs.Want1Zh + f.Hdb.Want1Zh + f.Msc.Want1Zh + f.Touch.Want1Zh + f.Xh.Want1Zh + f.Bj.Want1Zh + f.Sj.Want1Zh + f.Kfb.Want1Zh + f.Yb.Want1Zh +
		f.Bgs.Want2Zh + f.Hdb.Want2Zh + f.Msc.Want2Zh + f.Touch.Want2Zh + f.Xh.Want2Zh + f.Bj.Want2Zh + f.Sj.Want2Zh + f.Kfb.Want2Zh + f.Yb.Want2Zh
	f.TotalPF = f.Bgs.Want1Pf + f.Hdb.Want1Pf + f.Msc.Want1Pf + f.Touch.Want1Pf + f.Xh.Want1Pf + f.Bj.Want1Pf + f.Sj.Want1Pf + f.Kfb.Want1Pf + f.Yb.Want1Pf +
		f.Bgs.Want2Pf + f.Hdb.Want2Pf + f.Msc.Want2Pf + f.Touch.Want2Pf + f.Xh.Want2Pf + f.Bj.Want2Pf + f.Sj.Want2Pf + f.Kfb.Want2Pf + f.Yb.Want2Pf
	f.TotalMGS = f.Bgs.Want1Mgs + f.Hdb.Want1Mgs + f.Msc.Want1Mgs + f.Touch.Want1Mgs + f.Xh.Want1Mgs + f.Bj.Want1Mgs + f.Sj.Want1Mgs + f.Kfb.Want1Mgs + f.Yb.Want1Mgs +
		f.Bgs.Want2Mgs + f.Hdb.Want2Mgs + f.Msc.Want2Mgs + f.Touch.Want2Mgs + f.Xh.Want2Mgs + f.Bj.Want2Mgs + f.Sj.Want2Mgs + f.Kfb.Want2Mgs + f.Yb.Want2Mgs
	f.Total /= 2
	f.TotalToday /= 2
	f.TotalZH /= 2
	f.TotalPF /= 2
	f.TotalMGS /= 2

	utility.ResponseSuccess(c, gin.H{
		"data": f,
	})
}
