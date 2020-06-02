/**
 * @Author: alienongwlx@gmail.com
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020/6/2 14:02
 */

package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

func KeyGen(bits int, publicF, privateF string) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA Private key",
		Bytes: derStream,
	}
	privateFile, err := os.Create(privateF)
	defer privateFile.Close()
	err = pem.Encode(privateFile, block)
	if err != nil {
		return err
	}
	publicKey := &privateKey.PublicKey
	deRpKix, err := x509.MarshalPKIXPublicKey(publicKey)
	block = &pem.Block{
		Type:  "RSA Public key",
		Bytes: deRpKix,
	}
	if err != nil {
		return err
	}
	publicFile, err := os.Create(publicF)
	defer publicFile.Close()
	err = pem.Encode(publicFile, block)
	if err != nil {
		return err
	}
	return nil
}
func FileLoad(filePath string) []byte {
	privateFile, err := os.Open(filePath)
	defer privateFile.Close()
	if err != nil {
		return nil
	}
	privateKey := make([]byte, 2048)
	num, err := privateFile.Read(privateKey)
	return privateKey[:num]
}
func Encrypt(publicFile string, orgData []byte) ([]byte, error) {
	publicKey := FileLoad(publicFile)
	if publicKey == nil {
		return nil, errors.New("public file is bad")
	}
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key is bad")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, orgData) //加密
}
func Decrypt(privateFile string, cipherText []byte) ([]byte, error) {
	privateKey := FileLoad(privateFile)
	if privateKey == nil {
		return nil, errors.New("private file is bad")
	}
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("public key is bad")
	}
	p, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, p, cipherText)
}
