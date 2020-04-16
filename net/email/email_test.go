/**
 * @Author: alienongwlx@gmail.com
 * @Description: An Email Test Case
 * @Version: 1.0.0
 * @Date: 2020/4/16 14:05
 */

package email

import (
	"fmt"
	"testing"
)

func TestFetchMailboxInfo(t *testing.T){
	ms:=NewMailService("xxxxxxx:993","xxxxxxx@qq.com", "ndlwjilaoiaxbjhe")
	if err:=ms.Login();err!=nil{
		panic(err)
	}
	defer ms.Logout()
	mialboxs,err:=ms.FetchMailboxInfo()
	if err!=nil{
		panic(err)
	}
	fmt.Println(mialboxs)
	mc,err:=ms.FetchMailContent(2,"INBOX")
	if mc!=nil{

	}
}