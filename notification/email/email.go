/**
 * @Author: alienongwlx@gmail.com
 * @Description: send email by gomail
 * @Version: 1.0.0
 * @Date: 2020/4/15 19:50
 */
package email

import (
	gomail "gopkg.in/gomail.v2"
)

type MailService struct {
	Host     string
	Port     int
	Account  string
	Password string
}

type Email struct {
	To      []string
	Subject   string
	Content string
}

func NewMailService(Host string,Port int,Account,Password string)*MailService {
	return &MailService{Host,Port,Account,Password}
}

func (ms *MailService)SendMail(email *Email)error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(ms.Account, ms.Account)) // 添加别名
	m.SetHeader("To", email.To...)
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text/html", email.Content)
	d := gomail.NewDialer(ms.Host, ms.Port, ms.Account,ms.Password)
	err := d.DialAndSend(m)
	return err
}