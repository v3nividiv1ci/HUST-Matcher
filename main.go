package main

import (
	"HUST-Matcher/database"
	//"github.com/garyburd/redigo/redis"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	database.InitRClient()
	r := gin.Default()
	r = CollectRoute(r)
	//StudentID := "202013878"
	//auth.EmailAuth(StudentID)
	c := database.GetRClient()
	pong, err := c.Ping().Result()
	fmt.Println(pong, err)
	//err = http.ListenAndServe(":9990", nil)
	//if err != nil {
	//	return
	//}
	//fmt.Println(j)
	err = r.Run()
	if err != nil {
		panic(err)
	}

	//r := gin.Default()
	////r = CollectRoute(r)
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//err := r.Run()
	//if err != nil {
	//	return
	//}
}
