package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/png"
	"math/rand"
	"os"
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

// GetQRCode 生成二维码
func GetQRCode(serialId string) string {
	qrCode, _ := qr.Encode("http://localhost:8080/v1/asset/getAsset/"+serialId, qr.M, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 256, 256)
	file, _ := os.Create(fmt.Sprintf("qr_code_images/%s.png", serialId))
	defer file.Close()
	png.Encode(file, qrCode)
	return "qr_code_images/" + serialId + ".png"
}
