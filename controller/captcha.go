package controller

import (
	"fmt"
	"zjutjh/Join-Us/utility"

	"github.com/gin-gonic/gin"
)

func NewCaptcha(c *gin.Context) {
	id, b64s, err := utility.GetNewCaptcha()
	if err != nil {
		utility.ResponseError(c, err.Error())
		fmt.Println(err)
	}
	utility.ResponseSuccess(c, gin.H{
		"id":   id,
		"b64s": b64s,
	})
}

type VerifyCaptchaForm struct {
	id   string
	code string
}

func VerifyCaptcha(c *gin.Context) {
	var data VerifyCaptchaForm
	if err := c.ShouldBindJSON(&data); err != nil {
		utility.ResponseError(c, err.Error())
		return
	}
	if ok := utility.VerifyCaptcha(data.id, data.code); ok {
		utility.ResponseSuccess(c, gin.H{
			"status": "ok",
		})
	} else {
		utility.ResponseError(c, "验证码错误")
	}
}
