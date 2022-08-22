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
	if secret == Config.Secret || secret == Config.Secret2 {
		if secret == Config.Secret {
			c.Set("admin", "normal")
		} else {
			c.Set("admin", "advance")
		}
		c.Next()
	} else {
		utility.ResponseError(c, "No Authorization")
		c.Abort()
	}
}

func IsAdvanced(c *gin.Context) {
	secret := c.GetHeader("Authorization")
	if secret == "" {
		utility.ResponseError(c, "No Authorization")
		c.Abort()
		return
	}
	if secret == Config.Secret2 {
		c.Next()
	} else {
		utility.ResponseError(c, "No Authorization")
		c.Abort()
	}
}
