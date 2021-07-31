package auth

import (
	"HUST-Matcher/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JwtKey = []byte(model.KEY)

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func TokenRelease(user model.User) (string, error) {
	ExpTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: ExpTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "HUST-Matcher",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	TokenString, err := token.SignedString(JwtKey)

	if err != nil {
		return "", err
	}

	return TokenString, nil
}

func TokenParse(TokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(TokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	return token, claims, err

}
