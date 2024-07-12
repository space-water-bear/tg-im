package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GenerateToken(userId int64, username string, secret string, expire int64) (int64, string, error) {
	iat := time.Now().Unix()
	expTime := iat + expire*60*60*24
	claims := make(jwt.MapClaims)
	claims["id"] = userId
	claims["username"] = username
	claims["iat"] = iat
	claims["exp"] = expTime
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Claims = claims

	fmt.Println("gen token user:", userId, username)
	rest, err := token.SignedString([]byte(secret))
	return expTime, rest, err
}
