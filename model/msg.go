package model

import "gorm.io/gorm"

type Msg struct {
	gorm.Model
	StudentId      string `json:"studentID" `
	Title          string `json:"title" binding:"required"`
	Classification string `json:"classification" binding:"required"`
	Tag1           string `json:"tag1"`
	Tag2           string `json:"tag2"`
	Tag3           string `json:"tag3"`
	Time           string `json:"time" binding:"required"`
	Location       string `json:"location" binding:"required"`
	QQ             string `json:"qq" binding:"required"`
	Phone          string `json:"phone" binding:"required"`
	Date           string `json:"date" binding:"required"`
}

type Lost struct {
	gorm.Model
	StudentId      string `json:"studentID" `
	Title          string `json:"title" binding:"required"`
	Classification string `json:"classification" binding:"required"`
	Tag1           string `json:"tag1"`
	Tag2           string `json:"tag2"`
	Tag3           string `json:"tag3"`
	Time           string `json:"time" binding:"required"`
	Location       string `json:"location" binding:"required"`
	QQ             string `json:"qq" binding:"required"`
	Phone          string `json:"phone" binding:"required"`
	Date           string `json:"date" binding:"required"`
}

type Found struct {
	gorm.Model
	StudentId      string `json:"studentID"`
	Title          string `json:"title" binding:"required"`
	Classification string `json:"classification" binding:"required"`
	Tag1           string
	Tag2           string
	Tag3           string
	Time           string `json:"time" binding:"required"`
	Location       string `json:"location" binding:"required"`
	QQ             string `json:"qq" binding:"required"`
	Phone          string `json:"phone" binding:"required"`
	Date           string `json:"date" binding:"required"`
}
