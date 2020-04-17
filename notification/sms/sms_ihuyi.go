/**
 * @Author: alienongwlx@gmail.com
 * @Description: ihuyi sms interface
 * @Version: 1.0.0
 * @Date: 2020/4/17 9:02
 */

package sms

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

/**
@description: IHuYiSMSService Struct
@attribute Domain: IHuYi Url
@attribute Account: IHuYi Account
@attribute Password: IHuYi Password
*/
type IHuYiSMSService struct {
	Domain   string
	Account  string
	Password string
}

/**
@description: Submit Result Struct
@attribute Xmlns: Xmlns
@attribute Code: Code
@attribute Msg: Msg
@attribute SmsId: SmsId
*/
type SubmitResult struct {
	Xmlns string `xml:"xmlns,attr"`
	Code  int    `xml:"code"`
	Msg   string `xml:"msg"`
	SmsId string `xml:"smsid"`
}

/**
@description: NewIHuYiSMSService
@param account: IHuYi Account
@param password: IHuYi Password
@return: IHuYiSMSService
*/
func NewIHuYiSMSService(account, password string) ISMSService {
	return &IHuYiSMSService{
		Domain:   "http://106.ihuyi.cn/webservice/sms.php",
		Account:  account,
		Password: password,
	}
}

/**
@description: Send Message To Mobile
@param mobile: Mobile
@param message: Message
@param params: param
@return: error of nil
*/
func (s *IHuYiSMSService) SMS(mobile, message string, params map[string]string) error {
	req, err := url.Parse(s.Domain)
	if err != nil {
		return err
	}
	query := req.Query()
	query.Set("account", s.Account)
	query.Set("password", s.Password)
	query.Set("method", "Submit")
	query.Set("mobile", mobile)
	query.Set("content", message)
	req.RawQuery = query.Encode()
	fmt.Println(fmt.Sprintf("发送短信： %s - %s", mobile, message))
	resp, err := http.Get(req.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var result SubmitResult
	if err := xml.Unmarshal(bs, &result); err != nil {
		return err
	}
	if result.Code == 2 {
		return nil
	}
	return errors.New(result.Msg)
}
