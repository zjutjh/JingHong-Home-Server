package controller

import (
	"zjutjh/Join-Us/utility"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	utility.ResponseSuccess(c, nil)
}
