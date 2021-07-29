package main

import (
	"HUST-Matcher/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/user/auth/send", controller.RegisterAuth)

	return r
}
