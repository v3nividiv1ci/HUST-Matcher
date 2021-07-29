package auth

import (
	"HUST-Matcher/database"
	"HUST-Matcher/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get authorization header
		//oauth2.0 "bearer "[7]
		string := c.GetHeader("Authorization")
		if len(string) < 8 {
			c.JSON(http.StatusUnauthorized,
				gin.H{"code": 10010, "msg": "token格式错误"})
		}
		TString := c.GetHeader("Authorization")[7:]
		token, claims, err := TokenParse(TString)

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized,
				gin.H{"code": 10011, "msg": "token不合法"})
		}

		UserId := claims.UserId
		DB := database.GetDB()
		var user model.User
		DB.First(&user, UserId)

		// not registered
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized,
				gin.H{"code": 10012, "msg": "用户未注册"})
		}

		// write
		c.Set("user", user)
		c.Next()
	}
}
