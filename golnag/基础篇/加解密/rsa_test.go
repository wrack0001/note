package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnitDecryptAndEncrypt(t *testing.T) {
	privateKeyBytes, publicKeyBytes, _ := CreateKey(1024)
	fmt.Printf("privateKeyBytes = %s", privateKeyBytes)
	fmt.Printf("publicKeyBytes = %s", publicKeyBytes)
	text := []byte("1234")
	SecretText, _ := Encrypt(text, publicKeyBytes)
	fmt.Printf("SecretText = %s", SecretText)
	decryptText, _ := Decrypt(SecretText, privateKeyBytes)
	fmt.Printf("decryptText = %s", decryptText)
	assert.Equalf(t, text, decryptText, "TestUnitDecryptAndEncrypt")
}

// CreateKey 公钥、私钥生成器 参数bits: 指定生成的秘钥的长度, 单位: bit
func CreateKey(bits int) (privateKeyBytes, publicKeyBytes []byte, err error) {
	// 1. 生成私钥文件
	// GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	// 参数1: Reader是一个全局、共享的密码用强随机数生成器
	// 参数2: 秘钥的位数 - bit
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, fmt.Errorf("rsa GenerateKey")
	}
	// 2. MarshalPKCS1PrivateKey将rsa私钥序列化为ASN.1 PKCS#1 DER编码
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	// 3. Block代表PEM编码的结构, 对其进行设置
	block := pem.Block{
		Type:  "SPORTS RSA PRIVATE KEY",
		Bytes: derStream,
	}

	// 5. 使用pem编码, 并将数据写入文件中 (私钥)
	privateKeyBytes = pem.EncodeToMemory(&block)

	// 7. 生成公钥文件
	publicKey := privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return nil, nil, fmt.Errorf("rsa MarshalPKIXPublicKey")
	}
	block = pem.Block{
		Type:  "SPORTS RSA PUBLIC KEY",
		Bytes: derPkix,
	}

	// 8. 编码公钥, 写入文件
	publicKeyBytes = pem.EncodeToMemory(&block)
	return privateKeyBytes, publicKeyBytes, nil
}

// Encrypt rsa加密  text 要加密的数据 key 公钥
func Encrypt(text, key []byte) ([]byte, error) {
	// 1. 从数据中查找到下一个PEM格式的块
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, fmt.Errorf("encrypt pen Decode error is nil")
	}
	// 2. 解析一个DER编码的公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("x509 ParsePKIXPublicKey %v", err)
	}
	pubKey := pubInterface.(*rsa.PublicKey)

	// 3. 公钥加密
	result, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, text)
	if err != nil {
		return nil, fmt.Errorf("rsa EncryptPKCS1v15  %v", err)
	}
	return result, nil
}

// Decrypt rsa解密  text 要解密的数据 key 私钥文件的路径
func Decrypt(text, key []byte) ([]byte, error) {
	// 1. 从数据中查找到下一个PEM格式的块
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, fmt.Errorf("decrypt pen Decode error is nil")
	}
	// 2. 解析一个pem格式的私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("x509 ParsePKCS1PrivateKey %v", err)
	}
	// 3. 私钥解密
	result, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, text)
	if err != nil {
		return nil, fmt.Errorf("rsa DecryptPKCS1v15 %v", err)
	}
	return result, nil
}
