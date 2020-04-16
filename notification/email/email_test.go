/**
 * @Author: alienongwlx@gmail.com
 * @Description: email testing
 * @Version: 1.0.0
 * @Date: 2020/4/15 19:50
 */
package email

import "testing"

func TestSendMail(t *testing.T) {
	host:="smtp.sina.com"
	port:= 25
	Account:="xxxxx@sina.com"
	Password:="xxxxxxx"
	ms:=MailService{host,port,Account,Password}
	email:=Email{[]string{"xxxxx@qq.com"},"test","Test"}
	err:=ms.SendMail(&email)
	if err!=nil{
		panic(err)
	}
}