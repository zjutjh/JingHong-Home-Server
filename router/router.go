package router

import (
	"zjutjh/Join-Us/controller"
	"zjutjh/Join-Us/utility/middleware"
)

func SetRouter() {
	// TODO: routers list here
	Router.POST("/api/new_normal", middleware.VerifyCaptcha, controller.NewNormalForm)      // 招新季招新表单
	Router.POST("/api/new_develop", middleware.VerifyCaptcha, controller.NewDevelopForm)    // 长期招新：技术部 表单
	Router.GET("/api/get_forms/develop", middleware.IsAdmin, controller.GetAllDevelopForms) //获取所有招新季表单
	Router.GET("/api/get_forms/normal", middleware.IsAdmin, controller.GetAllNormalForms)   // 获取所有技术部表单
	// Router.GET("/api/get_forms/total", middleware.IsAdmin, controller.GetTotalForm)
	Router.GET("/api/get_forms/total", controller.GetFormsTotal) // 获取表单统计信息
	Router.GET("/api/get_captcha", controller.NewCaptcha)        // 获取验证码`
}
