package model

type Msg struct {
	//gorm.Model
	StudentId      string `json:"studentID" binding:"required"`
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

//type Seek struct {
//	StudentId  string `json:"studentID" binding:"required"`
//	Title  string `json:"title" binding:"required"`
//	Classification string `json:"classification" binding:"required"`
//	Tags1 string `json:"tags" binding:"required"`
//	Tags2 string `json:"tags" binding:"required"`
//	Tags3 string `json:"tags" binding:"required"`
//	Time string `json:"time" binding:"required"`
//	Location string `json:"location" binding:"required"`
//	QQ string `json:"qq" binding:"required"`
//	Phone string `json:"phone" binding:"required"`
//	Date string `json:"date" binding:"required"`
//}
