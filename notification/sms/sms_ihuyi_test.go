/**
 * @Author: alienongwlx@gmail.com
 * @Description: Ihuyi SMS Testing Case
 * @Version: 1.0.0
 * @Date: 2020/4/17 9:51
 */

package sms

import (
	"fmt"
	"testing"
)

func TestSMS(t *testing.T) {
	svs := NewIHuYiSMSService("cf_heda", "heda123!")
	var param map[string]string
	msg := fmt.Sprintf(`普通工单：工单编号【%s】,目前状态:【%s】。`, "测试工单号", "进行中")
	err := svs.SMS("17357325314", msg, param)
	if err != nil {
		t.Error(err)
	}
}
