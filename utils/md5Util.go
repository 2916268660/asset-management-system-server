package utils

import (
	"crypto/md5"
	"fmt"
)

var solt = "kjh1k2" //盐

// Encrypt md5加密
func Encrypt(password string) string {
	hash := md5.Sum([]byte(solt + "|" + password))
	return fmt.Sprintf("%x", hash)
}
