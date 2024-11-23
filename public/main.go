package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
)

func main() {
	// 创建一个密钥
	secretKey := []byte("your-secret-key")
	// 创建一个声明
	claims := jwt.MapClaims{
		"foo": "bar",
		"nbf": 15000,
	}
	// 创建一个新的 Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println("key", secretKey)
	fmt.Println("token", token)
	// 使用密钥签名 Token
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Fatalf("Error signing token: %v", err)
	}

	fmt.Println("Generated Token:", tokenString)
}
