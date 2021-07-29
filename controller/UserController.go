package controller

import (
	"HUST-Matcher/auth"
	"HUST-Matcher/database"
	"HUST-Matcher/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"regexp"
	"strings"
)

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello there",
	})
}

func AuthSend(c *gin.Context) {
	studentID := c.PostForm("studentID")
	// 验证学号格式
	match, _ := regexp.MatchString("^[UMDumd]\\d{9}$", studentID)
	if match != true {
		c.JSON(201,
			gin.H{"code": 10001, "msg": "学号格式错误"})
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
		gin.H{"code": 200, "msg": "验证码已发送", "auth": vcode})
}

func RegisterAuth(c *gin.Context) {
	studentID := c.PostForm("studentID")
	auth := c.PostForm("auth")
	match, _ := regexp.MatchString("^[UMDumd]\\d{9}$", studentID)
	if match != true {
		c.JSON(201,
			gin.H{"code": 10001, "msg": "学号格式错误"})
		return
	}
	// 如果已经注册，返回
	DB := database.GetDB()
	var user model.User
	DB.Table("users").Where("student_id = ?", studentID).First(&user)
	if user.ID != 0 {
		c.JSON(201,
			gin.H{"code": 10002, "msg": "学号已注册"})
		return
	}

	client := database.GetRClient()
	ID := strings.ToUpper(studentID)
	code, err := client.Get(ID).Result()
	if err != nil {
		c.JSON(201,
			gin.H{"code": 10003, "msg": "未获取验证码或验证码已过期"})
		return
	}
	if strings.Compare(auth, code) != 0 {
		c.JSON(201,
			gin.H{"code": 10004, "msg": "验证码错误"})
		return
	}
	// 验证成功删除验证码kv
	client.Del(ID)
	c.JSON(http.StatusOK,
		gin.H{"code": 200, "msg": "验证成功"})

}

func RegisterInfo(c *gin.Context) {
	DB := database.GetDB()
	var user model.User
	studentID := c.PostForm("studentID")
	username := c.PostForm("username")
	password := c.PostForm("password")
	// 虽然这个地方是前端把跳转之前保存的验证合法的学号传给我，但是还是想验证一下。。
	match, _ := regexp.MatchString("^[UMDumd]\\d{9}$", studentID)
	if match != true {
		c.JSON(201,
			gin.H{"code": 10001, "msg": "学号格式错误"})
		return
	}
	// 检验是否已注册
	DB.Table("users").Where("student_id = ?", studentID).First(&user)
	if user.ID != 0 {
		c.JSON(201,
			gin.H{"code": 10002, "msg": "学号已注册"})
		return
	}
	// 检验用户名、密码是否合法及长度
	if auth.CheckString(username) != true {
		c.JSON(201,
			gin.H{"code": 10005, "msg": "用户名不合法"})
		return
	}
	if auth.CheckString(password) != true {
		c.JSON(201,
			gin.H{"code": 10006, "msg": "密码不合法"})
		return
	}
	EncryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(201,
			gin.H{"code": 10007, "msg": "系统错误，密码加密中断"})
		return
	}

	User := model.User{
		Username:  username,
		Password:  string(EncryptedPassword),
		StudentId: studentID,
	}
	DB.Table("users").Create(&User)
	c.JSON(200,
		gin.H{"code": 200, "msg": "注册成功"})

}

func LoginPwd(c *gin.Context) {
	//	检查用户名密码是否合法
	DB := database.GetDB()
	var user model.User
	username := c.PostForm("username")
	password := c.PostForm("password")
	if auth.CheckString(username) != true {
		c.JSON(201,
			gin.H{"code": 10005, "msg": "用户名不合法"})
		return
	}
	if auth.CheckString(password) != true {
		c.JSON(201,
			gin.H{"code": 10006, "msg": "密码不合法"})
		return
	}
	DB.Table("users").Where("Username = ?", username).First(&user)
	if user.ID == 0 {
		c.JSON(201,
			gin.H{"code": 10008, "msg": "用户未注册"})
		return
	}
	//	解密数据库里的密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(201,
			gin.H{"code": 10009, "msg": "密码错误"})
		return
	}
	token, err := auth.TokenRelease(user)
	if err != nil {
		c.JSON(201,
			gin.H{"code": 10013, "msg": "token发放失败"})
	}
	c.JSON(http.StatusOK,
		gin.H{"code": 200, "data": gin.H{"token": token}, "msg": "登陆成功"})
}

func LoginEmail(c *gin.Context) {
	studentID := c.PostForm("studentID")
	auth_r := c.PostForm("auth")
	match, _ := regexp.MatchString("^[UMDumd]\\d{9}$", studentID)
	if match != true {
		c.JSON(201,
			gin.H{"code": 10001, "msg": "学号格式错误"})
		return
	}
	// 如果未注册
	DB := database.GetDB()
	var user model.User
	DB.Table("users").Where("student_id = ?", studentID).First(&user)
	if user.ID == 0 {
		c.JSON(201,
			gin.H{"code": 10008, "msg": "用户未注册"})
		return
	}

	// 检验验证码
	client := database.GetRClient()
	ID := strings.ToUpper(studentID)
	code, err := client.Get(ID).Result()
	if err != nil {
		c.JSON(201,
			gin.H{"code": 10003, "msg": "未获取验证码或验证码已过期"})
		return
	}
	if strings.Compare(auth_r, code) != 0 {
		c.JSON(201,
			gin.H{"code": 10004, "msg": "验证码错误"})
		return
	}

	// 验证成功删除验证码kv并发放token
	token, err := auth.TokenRelease(user)
	if err != nil {
		c.JSON(201,
			gin.H{"code": 10013, "msg": "token发放失败"})
	}
	client.Del(ID)
	c.JSON(http.StatusOK,
		gin.H{"code": 200, "data": gin.H{"token": token}, "msg": "验证成功"})
}
