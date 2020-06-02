/**
 * @Author: alienongwlx@gmail.com
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020/6/2 14:13
 */

package rsa

import (
	"flag"
	"testing"
)

func TestKeyGen(t *testing.T) {
	var bits int
	flag.IntVar(&bits, "b", 1024, "密码默认长度1024")
	err := KeyGen(bits, "public.pem", "private.pem")
	if err != nil {
		t.Error("KeyGen error" + err.Error())
	}
}
