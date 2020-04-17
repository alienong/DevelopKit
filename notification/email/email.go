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

/**
@description: MailService Struct
@attribute Host: EMail Server
@attribute Port: Port
@attribute Account: Email Account
@attribute Password: Email Password
*/
type MailService struct {
	Host     string
	Port     int
	Account  string
	Password string
}

/**
@description: Email Info Struct
@attribute To: EMail Receiver
@attribute Subject: Email Subject
@attribute Content: Email Content
*/
type Email struct {
	To      []string
	Subject string
	Content string
}

/**
@description: NewMailService
@param host: EMail Server
@param port: Port
@param account: EMail account
@param password: EMail password
@return: MailService
*/
func NewMailService(host string, port int, account, password string) *MailService {
	return &MailService{host, port, account, password}
}

/**
@description: SendMail
@param email: Email Struct Info
@return: error or nil
*/
func (ms *MailService) SendMail(email *Email) error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(ms.Account, ms.Account)) // 添加别名
	m.SetHeader("To", email.To...)
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text/html", email.Content)
	d := gomail.NewDialer(ms.Host, ms.Port, ms.Account, ms.Password)
	err := d.DialAndSend(m)
	return err
}
