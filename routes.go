package main

import (
	"HUST-Matcher/auth"
	"HUST-Matcher/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/user/auth/send", controller.AuthSend)
	r.POST("/api/user/register/auth", controller.RegisterAuth)
	r.POST("/api/user/register/info", controller.RegisterInfo)
	r.POST("/api/user/login/pwd", controller.LoginPwd)
	r.POST("/api/user/login/email", controller.LoginEmail)

	r.POST("/test/token", auth.JwtMiddleWare())
	r.POST("/api/seek/object", auth.JwtMiddleWare(), controller.SeekObject)
	r.GET("/test", controller.Test)
	return r
}
