package middleware

import (
	"strconv"
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
	id, err := utility.ParseToken(secret)
	if err != nil {
		utility.ResponseError(c, "Token Error")
		c.Abort()
	}
	id_int, ok := strconv.Atoi(id)
	if ok != nil {
		utility.ResponseError(c, "Token Error")
		c.Abort()
	}
	c.Set("id", uint64(id_int))
	c.Next()
}
