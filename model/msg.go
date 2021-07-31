package model

import (
	"gorm.io/gorm"
)

type Msg struct {
	gorm.Model
	StudentId      string `form:"studentID" `
	Title          string `form:"title" binding:"required"`
	Classification string `form:"classification" binding:"required"`
	Tag1           string `form:"tag1" binding:"required"`
	Tag2           string `form:"tag2" binding:"required"`
	Tag3           string `form:"tag3" binding:"required"`
	Detail         string `form:"detail" binding:"required"`
	Time           string `form:"time" binding:"required"`
	Location       string `form:"location" binding:"required"`
	QQ             string `form:"qq" binding:"required"`
	Phone          string `form:"phone" binding:"required"`
	Date           string `form:"date" binding:"required"`
}

type Lost struct {
	gorm.Model
	StudentId      string `form:"studentID" `
	Title          string `form:"title" binding:"required"`
	Classification string `form:"classification" binding:"required"`
	Tag1           string `form:"tag1" binding:"required"`
	Tag2           string `form:"tag2" binding:"required"`
	Tag3           string `form:"tag3" binding:"required"`
	Detail         string `form:"detail" binding:"required"`
	Time           string `form:"time" binding:"required"`
	Location       string `form:"location" binding:"required"`
	QQ             string `form:"qq" binding:"required"`
	Phone          string `form:"phone" binding:"required"`
	Date           string `form:"date" binding:"required"`
}

type Found struct {
	gorm.Model
	StudentId      string `form:"studentID"`
	Title          string `form:"title" binding:"required"`
	Classification string `form:"classification" binding:"required"`
	Tag1           string `form:"tag1" binding:"required"`
	Tag2           string `form:"tag2" binding:"required"`
	Tag3           string `form:"tag3" binding:"required"`
	Detail         string `form:"detail" binding:"required"`
	Time           string `form:"time" binding:"required"`
	Location       string `form:"location" binding:"required"`
	QQ             string `form:"qq" binding:"required"`
	Phone          string `form:"phone" binding:"required"`
	Date           string `form:"date" binding:"required"`
}
