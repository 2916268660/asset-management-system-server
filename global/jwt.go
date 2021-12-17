package global

import (
	"github.com/golang-jwt/jwt"
	"time"
)

const TokenExpireDuration = time.Hour * 12

var MySecret = []byte("a8x0sd.")

type MyClaims struct {
	StuId string `json:"stuId"`
	jwt.StandardClaims
}

// GetToken 生成token
func GetToken(stuId string) (string, error) {
	c := MyClaims{
		stuId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "root", //签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(MySecret)
}

func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
