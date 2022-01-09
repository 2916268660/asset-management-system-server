package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// GetValidateCode 生成六位随机验证码
func GetValidateCode() string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < 6; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

// md5盐
const solt = "kjh1k2"

// Encrypt md5加密
func Encrypt(password string) string {
	hash := md5.Sum([]byte(solt + "|" + password))
	return fmt.Sprintf("%x", hash)
}
