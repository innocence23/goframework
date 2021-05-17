package lib

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var RedisDB *redis.Client

func InitCache() {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     Config.RedisAddr,
		Password: Config.RedisPWD,
		DB:       Config.RedisDB,
	})

	if _, err := RedisDB.Ping(context.Background()).Result(); err != nil {
		log.Fatal("redis init error")
	}
}
