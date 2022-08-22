package controller

import (
	"zjutjh/Join-Us/utility"

	"github.com/gin-gonic/gin"
)

func TestAdmin(c *gin.Context) {
	user := c.GetString("admin")
	utility.ResponseSuccess(c, gin.H{"data": user})
}
