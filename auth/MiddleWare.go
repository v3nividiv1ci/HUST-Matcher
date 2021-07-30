package auth

import (
	"HUST-Matcher/database"
	"HUST-Matcher/model"
	"github.com/gin-gonic/gin"
)

func JwtMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get authorization header
		//oauth2.0 "bearer "[7]
		string := c.GetHeader("Authorization")
		if len(string) < 8 {
			c.JSON(201,
				gin.H{"code": 10010, "msg": "token格式错误"})
			c.Abort()
			return
		}
		TString := c.GetHeader("Authorization")[7:]
		token, claims, err := TokenParse(TString)

		if err != nil || !token.Valid {
			c.JSON(201,
				gin.H{"code": 10011, "msg": "token不合法"})
			c.Abort()
			return
		}

		UserId := claims.UserId
		DB := database.GetDB()
		var user model.User
		DB.First(&user, UserId)

		// not registered
		if user.ID == 0 {
			c.JSON(201,
				gin.H{"code": 10012, "msg": "用户未注册"})
			c.Abort()
			return
		}

		// write
		c.Set("user", user)
		c.Next()
		//c.JSON(http.StatusOK,
		//	gin.H{"code": 200, "msg": "验证成功"})
	}
}
