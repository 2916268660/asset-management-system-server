package utils

import (
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"server/global"
	"server/model"
	"time"
)

// GetToken 生成token
func GetToken(user *model.SysUser) (string, error) {
	c := model.MyClaims{
		user.UserId,
		user.UserName,
		user.Email,
		user.Phone,
		user.Department,
		user.Role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(global.TokenExpireDuration).Unix(),
			Issuer:    "root", //签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(global.MySecret)
}

func ParseToken(tokenString string) (*model.MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return global.MySecret, nil
	})
	if err != nil {
		global.GLOBAL_LOG.Error("解析token失败", zap.Error(err))
		return nil, err
	}
	if claims, ok := token.Claims.(*model.MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
