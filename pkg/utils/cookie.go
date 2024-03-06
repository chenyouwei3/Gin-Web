package utils

import (
	"fmt"
	"github.com/gorilla/securecookie"
	"github.com/sirupsen/logrus"
)

var CookieKey = []byte("cyw")

func CookieEncryption(name, value string) string {
	// 实例化 securecookie
	var secure = securecookie.New(CookieKey, nil)
	// 对value进行编码
	encodeValue, err := secure.Encode(name, value)
	if err != nil {
		logrus.Error("cookie加密失败:", err)
	}
	return encodeValue
}

func CookieDecrypt(name, encodeValue string) string {
	// 实例化 securecookie
	var secure = securecookie.New(CookieKey, nil)
	// 对value进行编码
	var value string
	secure.Decode(name, encodeValue, &value)
	fmt.Println("value:", value)
	return value
}
