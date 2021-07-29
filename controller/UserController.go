package controller

import (
	"HUST-Matcher/auth"
	"HUST-Matcher/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strings"
)

func RegisterAuth(c *gin.Context) {
	studentID := c.PostForm("studentID")
	match, _ := regexp.MatchString("^[UMDumd]\\d{9}$", studentID)
	if match != true {
		c.JSON(http.StatusBadRequest,
			gin.H{"code": 400, "msg": "学号格式错误"})
		return
	}
	//	先查找，如果有，则删除
	client := database.GetRClient()
	ID := strings.ToUpper(studentID)
	_, err := client.Get(ID).Result()
	if err == nil {
		client.Del(ID)
	}
	vcode := auth.EmailAuth(strings.ToUpper(studentID))
	c.JSON(http.StatusOK,
		gin.H{"code": 200, "msg": "验证码已发送", "vericode": vcode})
}
