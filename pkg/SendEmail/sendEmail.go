package SendEmail

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"go-talk-talk/config"
	"net/smtp"
)

func SendEmail(sendTo []string, topic string, content []byte) error {
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = config.EmailConf.From
	// 设置接收方的邮箱
	e.To = sendTo
	//设置主题
	e.Subject = topic
	//设置文件发送的内容
	e.Text = content
	//设置服务器相关的配置
	err := e.SendWithTLS(
		config.EmailConf.SmtpAddr,
		smtp.PlainAuth("", config.EmailConf.SmtpUsername, config.EmailConf.SmtpPassword, config.EmailConf.SmtpHost),
		&tls.Config{
			ServerName: config.EmailConf.SmtpHost,
		},
	)
	return err
}
