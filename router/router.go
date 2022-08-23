package router

import (
	"zjutjh/Join-Us/controller"
	"zjutjh/Join-Us/utility/middleware"
)

func SetRouter() {
	Router.POST("/api/new_normal", middleware.VerifyCaptcha, controller.NewNormalForm)    // 招新季招新表单
	Router.GET("/api/get_forms/total", middleware.IsAdmin, controller.GetNormalFormTotal) // 获取表单统计信息
	Router.GET("/api/get_captcha", controller.NewCaptcha)                                 // 获取验证码
	Router.GET("/api/test_admin", middleware.IsAdmin, controller.TestAdmin)
	Router.GET("/api/all_normal_forms_brief", middleware.IsAdvanced, controller.GetAllNormalFormsBrief)
	Router.GET("/api/normal_form", middleware.IsAdvanced, controller.GetNormalForm)
	Router.GET("/api/export_normal_form", middleware.IsAdvanced, controller.ExportAllNormalFormExcel)
	Router.GET("/api/test", controller.Test)
}
