package utils

import (
	"errors"
	"github.com/go-ini/ini"
	"go.uber.org/zap"
	"gopkg.in/mail.v2"
	"server/global"
)

var (
	from string
	port int
	host string
	auth string //qq授权码
)

func init() {
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		global.GLOBAL_LOG.Error("加载配置文件失败", zap.Error(err))
		return
	}
	emailCfg := cfg.Section("email")
	from = emailCfg.Key("from").MustString("")
	port = emailCfg.Key("port").MustInt(25)
	host = emailCfg.Key("host").MustString("")
	auth = emailCfg.Key("auth").MustString("")
}

type EmailRequest struct {
	Emails []string // 请求发送的邮箱集合
	Title  string   // 邮箱标题
	Body   string   // 邮箱正文
}

// SendEmail 发送邮件
func (e *EmailRequest) SendEmail() error {

	message := mail.NewMessage()
	message.SetHeaders(
		map[string][]string{
			"From": []string{
				from,
			},
			"To": e.Emails,
			"Subject": []string{
				e.Title,
			},
		})
	message.SetBody("text/html", e.Body)

	dialer := mail.NewDialer(host, port, from, auth)
	err := dialer.DialAndSend(message)
	if err != nil {
		global.GLOBAL_LOG.Error("发送邮件失败", zap.Error(err))
		return errors.New("发送邮件失败")
	}
	return nil
}
