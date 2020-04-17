/**
 * @Author: alienongwlx@gmail.com
 * @Description: Http Get Methods
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
@description: Just Url Get Request
@param url: url
@return: result and error
*/
func HttpGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

/**
@description: Http Get Request Has Cookie
@param url: url
@param cookie: cookie string
@param param: param
@return: result and error
*/
func HttpGetHasCookie(url, cookie, param string) (string, []string, error) {
	rCookie := make([]string, 0)
	client := &http.Client{}
	parasByte := []byte(param)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(parasByte))
	if err != nil {
		return "", rCookie, err
	}
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
