package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	SecretKey = "243223ffslsfsldfl412fdsfsdf"
	ExpTime   = time.Minute * 60
)

func CreateToken(uid int) (string, error) {
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
