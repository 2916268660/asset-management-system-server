package global

import (
	"github.com/golang-jwt/jwt"
	"server/models/common"
	"time"
)

// GetToken 生成token
func GetToken(user *common.User) (string, error) {
	c := common.MyClaims{
		user.UserId,
		user.UserName,
		user.Email,
		user.Phone,
		user.Department,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "root", //签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(MySecret)
}

func ParseToken(tokenString string) (*common.MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &common.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*common.MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
