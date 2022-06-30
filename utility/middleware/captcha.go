package middleware

import (
	"zjutjh/Join-Us/utility"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type captchaStruct struct {
	Code string `json:"captcha_code"`
	Id   string `json:"captcha_id"`
}

func VerifyCaptcha(c *gin.Context) {
	var data captchaStruct
	if err := c.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		utility.ResponseError(c, err.Error())
		c.Abort()
		return
	}
	if ok := utility.VerifyCaptcha(data.Id, data.Code); ok {
		c.Next()
	} else {
		utility.ResponseError(c, "验证码错误")
		c.Abort()
	}
}
