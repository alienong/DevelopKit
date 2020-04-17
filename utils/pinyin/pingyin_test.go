/**
 * @Author: alienongwlx@gmail.com
 * @Description: A Pinyin Testing case
 * @Version: 1.0.0
 * @Date: 2020/4/17 11:37
 */

package pinyin

import (
	"fmt"
	"testing"
)

func TestGoChineseToPinyin(t *testing.T) {
	res, err := GoChineseToPinyin("我是谁", false)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)
}
