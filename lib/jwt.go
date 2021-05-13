package lib

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(uid int) (string, error) {
	SecretKey := Config.TokenSecretKey
	ExpTime := Config.TokenDuration
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  uid,
		"exp": time.Now().Add(ExpTime).Unix(),
	})
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	SecretKey := Config.TokenSecretKey
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
