package lib

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/cast"
)

var (
	ctx                   = context.Background()
	accessReidsKey        = "access_token_%s"
	refreshAccessReidsKey = "refresh_access_token_%s"
)

func CreateToken(uid int) (map[string]string, error) {
	SecretKey := Config.TokenSecretKey
	ExpTime := Config.TokenDuration
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  uid,
		"exp": time.Now().Add(ExpTime).Unix(),
	})
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return nil, err
	}
	refreshTokenString, err := createRefreshToken(uid)
	if err != nil {
		return nil, err
	}
	if err := CreateAuth(uid); err != nil {
		return nil, err
	}
	tokens := map[string]string{
		"access_token":  tokenString,
		"refresh_token": refreshTokenString,
	}

	return tokens, nil
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

func createRefreshToken(uid int) (string, error) {
	RefreshSecretKey := Config.RefreshTokenSecretKey
	ExpTime := Config.TokenDuration
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  uid,
		"exp": time.Now().Add(ExpTime).Unix(),
	})
	refreshTokenString, err := token.SignedString([]byte(RefreshSecretKey))
	if err != nil {
		return "", err
	}
	return refreshTokenString, nil
}

func parseRefreshToken(refreshToken string) (jwt.MapClaims, error) {
	RefreshSecretKey := Config.RefreshTokenSecretKey
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(RefreshSecretKey), nil
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

func RefreshToken(refreshToken string) (map[string]string, error) {
	claims, err := parseRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}
	uid := cast.ToInt(claims["id"])
	if uid == 0 {
		err = errors.New("token data conv error")
		return nil, err
	}
	if _, err := DeleteAuth(uid); err != nil {
		return nil, err
	}
	tokens, err := CreateToken(uid)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func CreateAuth(uid int) error {
	ExpTime := Config.TokenDuration
	errAccess := RedisDB.Set(ctx, genJWTRTRedisKey(uid), uid, ExpTime).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := RedisDB.Set(ctx, genJWTATRedisKey(uid), uid, ExpTime).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

func FetchAuth(uid int) (int, error) {
	result, err := RedisDB.Get(ctx, genJWTRTRedisKey(uid)).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := cast.ToIntE(result)
	return userID, nil
}

func DeleteAuth(uid int) (int64, error) {
	deleted, err := RedisDB.Del(ctx, genJWTRTRedisKey(uid)).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

func genJWTRTRedisKey(uid int) string {
	return fmt.Sprintf(accessReidsKey, cast.ToString(uid))
}

func genJWTATRedisKey(uid int) string {
	return fmt.Sprintf(refreshAccessReidsKey, cast.ToString(uid))
}
