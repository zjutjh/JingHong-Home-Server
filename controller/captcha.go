package controller

import (
	"log"
	"zjutjh/Join-Us/utility"

	"github.com/gin-gonic/gin"
)

func NewCaptcha(c *gin.Context) {
	id, b64s, err := utility.GetNewCaptcha()
	if err != nil {
		utility.ResponseError(c, err.Error())
		log.Println(err)
	}
	utility.ResponseSuccess(c, gin.H{
		"id":   id,
		"b64s": b64s,
	})
}
