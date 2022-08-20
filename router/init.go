package router

import (
	. "zjutjh/Join-Us/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	if !Config.Dev {
		gin.SetMode(gin.ReleaseMode)
	}
	Router = gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	if Config.Dev {
		corsConfig.AllowAllOrigins = true
	} else {
		corsConfig.AllowOrigins = Config.Server.AllowOrigins
	}
	Router.Use(cors.New(corsConfig))
	SetRouter()
}
