package database

import (
	"HUST-Matcher/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = model.USERNAME
const DB_PASSWORD = model.PASSWORD
const DB_NAME = "HM"
const DB_HOST = "123.56.0.167"
const DB_PORT = "3306"

var MDB *gorm.DB

func InitDB() *gorm.DB {
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Msg{})
	db.AutoMigrate(&model.Lost{})
	db.AutoMigrate(&model.Found{})
	MDB = db
	return db
}

func GetDB() *gorm.DB {
	return MDB
}
