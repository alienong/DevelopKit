/**
 * @Author: alienongwlx@gmail.com
 * @Description: Sms Service
 * @Version: 1.0.0
 * @Date: 2020/4/17 9:39
 */

package sms

/**
@description: SMSService Interface
@method SMS: Send Messages To Mobile
*/
type ISMSService interface {
	SMS(mobile, message string, params map[string]string) error
}
