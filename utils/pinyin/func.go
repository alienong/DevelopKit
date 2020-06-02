/**
 * @Author: alienongwlx@gmail.com
 * @Description: A Pinyin Translate
 * @Version: 1.0.0
 * @Date: 2020/4/15 19:50
 */

package pinyin

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/Tang-RoseChild/mahonia"
	"strings"
)

/**
@description: Chinese To Pinyin
@param chinese: chinese string
@return: pinyin string,error
*/
func ChineseToPinyin(chinese string) (string, error) {
	return GoChineseToPinyin(chinese, true)
}

/**
@description: Chinese To Pinyin
@param chinese: chinese string
@param full: Full Pinyin or Simple Pinyin
@return: pinyin string,error
*/
func GoChineseToPinyin(chinese string, full bool) (string, error) {
	enc := mahonia.NewEncoder("GBK")
	str := enc.ConvertString(chinese)
	var retStr string
	strHex := hex.EncodeToString([]byte(str))
	for index := 0; index <= len(strHex)-2; {
		var sLeft, sRight string
		sLeft = strHex[index : index+2]
		if index+4 <= len(strHex) {
			sRight = strHex[index+2 : index+4]
		} else {
			sRight = ""
		}
		bLeft, err := hex.DecodeString(sLeft)
		if err != nil {
			return "", errors.New("hex转换错误:" + err.Error())
		}
		bright, err := hex.DecodeString(sRight)
		if err != nil {
			return "", errors.New("hex转换错误:" + err.Error())
		}
		ileft := int(bLeft[0])
		var iright int
		if len(bright) > 0 {
			iright = int(bright[0])
		} else {
			iright = 0
		}
		var myCode int
		if ileft > 128 {
			index = index + 4
			myCode = 65536 - ileft*256 - iright
			for {
				if myCode > 25358 {
					fmt.Printf("%d\n", myCode)
					retStr = retStr + "a"
					break
				}
				if str, ok := CRT[myCode]; ok {
					if full {
						retStr = retStr + str
					} else {
						retStr = retStr + string(str[0])
					}
					break
				}
				myCode = myCode + 1
			}
		} else {
			if full {
				retStr = retStr + string([]byte{byte(ileft)})
			} else {
				retStr = retStr + string([]byte{byte(ileft)}[0])
			}
			index = index + 2
		}
	}
	return strings.ToUpper(retStr), nil
}
