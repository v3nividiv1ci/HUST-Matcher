package algorithm

import (
	"HUST-Matcher/database"
	"HUST-Matcher/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

//求并集长度
func union(slice1, slice2 [3]string) int {
	m := make(map[string]int)
	for _, v := range slice1 {
		m[v]++
	}
	ul := 3
	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			ul++
		}
	}
	return ul
}

//求交集长度
func intersect(slice1, slice2 [3]string) int {
	m := make(map[string]int)
	il := 0
	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			il++
		}
	}
	return il
}

// Match Pugss: 失魂人(去匹配的) Monee: 控魂人(被匹配到的) :)
func Match(c *gin.Context) {
	mid, _ := strconv.Atoi(c.PostForm("mid"))
	fmt.Println("mid was", mid)
	DB := database.GetDB()
	var Pugss model.Msg
	var Monees []model.Msg
	var l [2]string
	var kwP, kwM [3]string
	var class int                // 为何种类型
	var similarity, temp float32 // 相似度，由len(交集)/len(并集)求出
	tid := 0
	if mid/20000 == 1 {
		mid = mid % 20000
		class = 2
		l[0], l[1] = "founds", "losts"
	} else {
		mid = mid % 10000
		class = 1
		l[0], l[1] = "losts", "founds"
	}
	fmt.Println("class 是", class)
	fmt.Println("mid is", mid)
	TableP := DB.Table(l[0])
	TableP.Where("id = ?", mid).First(&Pugss)
	kwP[0], kwP[1], kwP[2] = Pugss.Tag2, Pugss.Tag1, Pugss.Tag3
	//	遍历
	//DB.Table(l[1]).Last(&Monee)
	result := DB.Table(l[1]).Find(&Monees)
	fmt.Println(Monees)
	num := int(result.RowsAffected)
	//num := int(Monee.ID)
	fmt.Println("num is", num)
	TimeP, _ := strconv.Atoi(Pugss.Time[0:4] + Pugss.Time[5:7] + Pugss.Time[8:10] + string(Pugss.Time[11]))
	for i := 0; i < num; i++ {
		fmt.Println("i ", i)
		//db := database.GetDB()
		//db.Table(l[1]).Where("id = ?", i).Take(&Monee)
		//if Monees[i].DeletedAt  {
		//	continue
		//}
		fmt.Println("monee id是", Monees[i].ID)
		//	1.时间里的数字 比大小
		TimeM, _ := strconv.Atoi(Monees[i].Time[0:4] + Monees[i].Time[5:7] + Monees[i].Time[8:10] + string(Monees[i].Time[11]))
		fmt.Println("时间是：", TimeP, TimeM)
		if (class == 2 && TimeP < TimeM) || (class == 1 && TimeP > TimeM) {
			continue
		}
		//	2.取出待匹配的中的三个关键词，进行匹配
		kwM[0], kwM[1], kwM[2] = Monees[i].Tag1, Monees[i].Tag2, Monees[i].Tag3
		temp = float32(intersect(kwM, kwP)) / float32(union(kwM, kwP))
		fmt.Println("%s 的匹配度是%s", i, temp)
		if temp >= similarity {
			similarity = temp
			tid = i
		}
	}
	fmt.Println("tid 是", tid)
	fmt.Println("monees[tid].id是", Monees[tid].ID)
	//dB := database.GetDB()
	//dB.Table(l[1]).Where("id = ?", tid).First(&Monee)
	if Monees[tid].ID == 0 || similarity == 0 {
		c.JSON(201,
			gin.H{"code": 10014, "msg": "无匹配", "post": ""})
		c.Abort()
		return
	}
	c.JSON(200,
		gin.H{"code": 200, "msg": "有匹配", "post": Monees[tid]})

}
