/**
 * @Author: alienongwlx@gmail.com
 * @Description: sha crypto
 * @Version: 1.0.0
 * @Date: 2020/4/17 14:32
 */

package sha256

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
)

/**
@description: GetRandomStringFunc Base ob tp
@param tp: 0 Means String;1 Means Character
@return: func(int) string
*/
func GetRandomStringFunc(tp int) func(int) string {
	bytes := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if tp == 1 {
		bytes = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%&*_+")
	}
	return func(length int) string {
		result := make([]byte, length)
		for i := 0; i < length; i++ {
			result[i] = bytes[rand.Intn(len(bytes))]
		}
		return string(result)
	}
}

/**
@description: GetRandomString
@param length: length
@return: string
*/
func GetRandomString(length int) string {
	return GetRandomStringFunc(0)(length)
}

/**
@description: GetRandomCharacter
@param length: length
@return: string
*/
func GetRandomCharacter(length int) string {
	return GetRandomStringFunc(1)(length)
}

/**
@description: SHA Cypto
@param src: Original String
@param salt: Salt String
@return: string
*/
func SHAEncrypt(src, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(src))
	hash.Write([]byte("$"))
	hash.Write([]byte(salt))
	return hex.EncodeToString(hash.Sum(nil))
}
