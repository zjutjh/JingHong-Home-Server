package controller

import (
	"zjutjh/Join-Us/utility"

	"github.com/gin-gonic/gin"
)

func TestAdmin(c *gin.Context) {
	utility.ResponseSuccess(c, nil)
}
