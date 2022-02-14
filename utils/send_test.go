package utils

import (
	"fmt"
	"server/model"
	"testing"
)

func TestSendEmail(t *testing.T) {
	code := GetValidateCode()
	msg := &EmailRequest{
		Emails: []string{
			"2916268660@qq.com",
			"1821064662@qq.com",
			"b.291808981@foxmail.com",
		},
		Title: model.EmailTitle_ValidateCode,
		Body:  fmt.Sprintf(model.EmailBody_validateCode, code),
	}
	err := msg.SendEmail()
	if err != nil {
		t.Error(err)
	}
}
