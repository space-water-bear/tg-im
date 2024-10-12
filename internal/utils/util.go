package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// Encrypt 加密, 返回加密后的字符串.
func Encrypt(source string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

// Compare 比较, 返回nil表示密码匹配.
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// TernaryIF 三元运算符
func TernaryIF[T any](condition bool, a, b T) T {
	if condition {
		return a
	}
	return b
}

// CompressMessage 压缩消息
func CompressMessage(message any) (string, error) {
	// 实现压缩算法
	jsonData, err := json.Marshal(message)
	if err != nil {
		fmt.Println("error json marshalling message", err)
		return "", err
	}

	var compressedData bytes.Buffer
	gz := gzip.NewWriter(&compressedData)
	if _, err := gz.Write(jsonData); err != nil {
		fmt.Println("error gzip compressing message", err)
		return "", err
	}
	if err := gz.Close(); err != nil {
		fmt.Println("error gzip close writer", err)
	}

	encodedData := base64.StdEncoding.EncodeToString(compressedData.Bytes())

	return encodedData, nil
}
