package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GenerateToken(userId int64, username string, secret string, expire int64) (string, error) {
	iat := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["id"] = userId
	claims["username"] = username
	claims["iat"] = iat
	claims["exp"] = iat + expire
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Claims = claims
	return token.SignedString([]byte(secret))
}
