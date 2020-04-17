/**
 * @Author: alienongwlx@gmail.com
 * @Description:Http Post Methods
 * @Version: 1.0.0
 * @Date: 2020/4/17 11:47
 */

package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

/**
@description: Http Post Request Has Cookie
@param url: url
@param cookie: cookie string
@param param: param
@param contentType: Content-Type e.g. application/json,application/x-www-form-urlencoded
@return: result and cookie and error
*/
func HttpPost(url, cookie, param, contentType string) (string, []string, error) {
	rCookie := make([]string, 0)
	client := &http.Client{}
	paramByte := []byte(param)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(paramByte))
	if err != nil {
		return "", rCookie, err
	}
	req.Header.Set("Content-Type", contentType)
	if len(cookie) > 0 {
		req.Header.Set("Cookie", cookie)
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", rCookie, err
	}
	defer resp.Body.Close()
	if _, ok := resp.Header["Set-Cookie"]; ok {
		rCookie = resp.Header["Set-Cookie"]
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", rCookie, err
	}
	return string(body), rCookie, nil
}
