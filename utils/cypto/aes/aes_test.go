/**
 * @Author: alienongwlx@gmail.com
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020/6/2 14:23
 */

package aes

import (
	"fmt"
	"testing"
)

func TestDecrypt(t *testing.T) {
	x := []byte("世界上最邪恶最专制的现代奴隶制国家--朝鲜")
	key := []byte("hgfedcba87654321")
	fmt.Println("加密前:", string(x))
	x1, err := Encrypt(x, key)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("加密后:", string(x1))
	x2, err := Decrypt(x1, key)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("解密后:", string(x2))
}
