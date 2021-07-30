package controller

import (
	"HUST-Matcher/database"
	"HUST-Matcher/model"
	"encoding/json"
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
	err := c.ShouldBindBodyWith(&lost, binding.JSON)
	if err != nil {
		fmt.Println(err)
		return
	}
	rawData, _ := c.Get(gin.BodyBytesKey)
	var f interface{}
	json.Unmarshal(rawData.([]byte), &f)
	tags := f.(map[string]interface{})["tags"].([]interface{})
	lost.Tag1, lost.Tag2, lost.Tag3 = tags[0].(string), tags[1].(string), tags[2].(string)
	lost.StudentId = studentID

	DB := database.GetDB()
	DB.Table("losts").Create(&lost)
	//DB.Table("losts").Where("student_id = ?", studentID).Last(&lost2)
	mid := lost.ID + 10000
	fmt.Println("the msg is", lost)
	c.JSON(200,
		gin.H{"code": 200, "msg": "发布成功", "mid": mid})
}

// mid: 2xxxx
func SeekPerson(c *gin.Context) {
	u, _ := c.Get("user")
	studentID := u.(model.User).StudentId
	// 拿从中间件里面解析出来的id
	var found model.Found
	err := c.ShouldBindBodyWith(&found, binding.JSON)
	if err != nil {
		fmt.Println(err)
		return
	}
	rawData, _ := c.Get(gin.BodyBytesKey)
	var f interface{}
	json.Unmarshal(rawData.([]byte), &f)
	tags := f.(map[string]interface{})["tags"].([]interface{})
	found.Tag1, found.Tag2, found.Tag3 = tags[0].(string), tags[1].(string), tags[2].(string)
	found.StudentId = studentID

	//var found2 model.Found
	DB := database.GetDB()
	DB.Table("founds").Create(&found)
	//DB.Table("founds").Where("student_id = ?", studentID).Last(&found)
	mid := found.ID + 20000
	fmt.Println("the msg is", found)
	c.JSON(200,
		gin.H{"code": 200, "msg": "发布成功", "mid": mid})
}
