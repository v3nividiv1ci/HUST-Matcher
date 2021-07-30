package controller

import (
	"HUST-Matcher/database"
	"HUST-Matcher/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func SeekObject(c *gin.Context) {
	var msg model.Msg
	err := c.ShouldBindBodyWith(&msg, binding.JSON)
	if err != nil {
		fmt.Println(err)
		return
	}
	rawData, _ := c.Get(gin.BodyBytesKey)
	var f interface{}
	json.Unmarshal(rawData.([]byte), &f)
	tags := f.(map[string]interface{})["tags"].([]interface{})
	msg.Tag1, msg.Tag2, msg.Tag3 = tags[0].(string), tags[1].(string), tags[2].(string)

	DB := database.GetDB()
	DB.Table("msgs").Create(&msg)

	fmt.Println("the msg is", msg)
	c.JSON(200,
		gin.H{"code": 200, "msg": "发布成功"})
}

//func SeekPerson(c *gin.Context) {
//	var msg model
//}
