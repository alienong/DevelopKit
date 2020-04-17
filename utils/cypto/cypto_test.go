/**
 * @Author: alienongwlx@gmail.com
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020/4/17 15:13
 */

package cypto

import (
	"fmt"
	"testing"
)

func TestSHA(t *testing.T) {
	s1 := GetRandomString(10)
	fmt.Println(s1)
	s2 := GetRandomCharacter(10)
	fmt.Println(s2)
	s3 := SHAEncrypt("alien", "hu")
	fmt.Println(s3)
}
