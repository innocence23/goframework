package lib

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Env struct {
	DBDriver              string        `mapstructure:"DB_DRIVER"`
	DBSource              string        `mapstructure:"DB_SOURCE"`
	ServerAddress         string        `mapstructure:"SERVER_ADDRESS"`
	TokenSecretKey        string        `mapstructure:"TOKEN_SECRET_KEY"`
	RefreshTokenSecretKey string        `mapstructure:"REFRESH_TOKEN_SECRET_KEY"`
	TokenDuration         time.Duration `mapstructure:"TOKEN_DURATION"`

	RedisAddr string `mapstructure:"REDIS_ADDR"`
	RedisPWD  string `mapstructure:"REDIS_PWD"`
	RedisDB   int    `mapstructure:"REDIS_DB"`
}

var Config Env

func LoadConfig(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("init env error")
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatal("init env error")
	}
}
