package controller

import (
	"HUST-Matcher/database"
	"HUST-Matcher/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// mid: 1xxxx
func SeekObject(c *gin.Context) {
	u, _ := c.Get("user")
	studentID := u.(model.User).StudentId
	// 拿从中间件里面解析出来的id
	var lost model.Lost
	err := c.BindWith(&lost, binding.Form)
	if err != nil {
		c.JSON(201,
			gin.H{"code": 10015, "msg": "字段不能为空"})
		c.Abort()
		return
	}
	//rawData, _ := c.Get(gin.BodyBytesKey)
	//var f interface{}
	//err = json.Unmarshal(rawData.([]byte), &f)
	//if err != nil {
	//	c.JSON(201,
	//		gin.H{"code": 10015, "msg": "字段不能为空"})
	//	c.Abort()
	//	return
	//}
	//tags := f.(map[string]interface{})["tags"].([]interface{})
	//lost.Tag1, lost.Tag2, lost.Tag3 = tags[0].(string), tags[1].(string), tags[2].(string)
	lost.StudentId = studentID

	DB := database.GetDB()
	DB.Table("losts").Create(&lost)
	//DB.Table("losts").Where("student_id = ?", studentID).Last(&lost2)
	mid := lost.ID + 10000
	fmt.Println("the msg is", lost)
	c.JSON(200,
		gin.H{"code": 200, "msg": "发布成功", "data": gin.H{"mid": mid}})
}

// mid: 2xxxx
func SeekPerson(c *gin.Context) {
	u, _ := c.Get("user")
	studentID := u.(model.User).StudentId
	// 拿从中间件里面解析出来的id
	var found model.Found
	err := c.BindWith(&found, binding.Form)
	if err != nil {
		c.JSON(201,
			gin.H{"code": 10015, "msg": "字段不能为空"})
		c.Abort()
		return
	}
	//rawData, _ := c.Get(gin.BodyBytesKey)
	//var f interface{}
	//err := json.Unmarshal(rawData.([]byte), &f)
	//if err != nil {
	//	return
	//}
	//tags := f.(map[string]interface{})["tags"].([]interface{})
	//found.Tag1, found.Tag2, found.Tag3 = tags[0].(string), tags[1].(string), tags[2].(string)
	found.StudentId = studentID

	//var found2 model.Found
	DB := database.GetDB()
	DB.Table("founds").Create(&found)
	//DB.Table("founds").Where("student_id = ?", studentID).Last(&found)
	mid := found.ID + 20000
	fmt.Println("the msg is", found)
	c.JSON(200,
		gin.H{"code": 200, "msg": "发布成功", "data": gin.H{"mid": mid}})
}

func MyPosts(c *gin.Context) {
	u, _ := c.Get("user")
	studentID := u.(model.User).StudentId
	DB := database.GetDB()
	var losts []model.Msg
	DB.Table("losts").Where("student_id = ?", studentID).Find(&losts)
	var founds []model.Msg
	DB.Table("founds").Where("student_id = ?", studentID).Find(&founds)
	var msgs []model.Msg
	msgs = append(losts, founds...)
	c.JSON(200,
		gin.H{"code": 200, "msg": "获取成功", "data": msgs})
}
