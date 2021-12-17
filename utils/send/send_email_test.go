package send

import (
	"fmt"
	"server/models/common"
	"server/utils"
	"testing"
)

func TestSendEmail(t *testing.T) {
	code := utils.GetValidateCode()
	msg := &EmailRequest{
		Emails: []string{
			"2916268660@qq.com",
			"1821064662@qq.com",
			"b.291808981@foxmail.com",
		},
		Title: common.EmailTitle_ValidateCode,
		Body:  fmt.Sprintf(common.EmailBody_validateCode, code),
	}
	err := msg.SendEmail()
	if err != nil {
		t.Error(err)
	}
}
