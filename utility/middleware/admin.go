package middleware

import (
	. "zjutjh/Join-Us/config"
	"zjutjh/Join-Us/utility"

	"github.com/gin-gonic/gin"
)

func IsAdmin(c *gin.Context) {
	secret := c.GetHeader("Authorization")
	if secret == "" {
		utility.ResponseError(c, "No Authorization")
		c.Abort()
		return
	}
	if secret == Config.Secret {
		c.Next()
	} else {
		utility.ResponseError(c, "No Authorization")
		c.Abort()
	}
}
