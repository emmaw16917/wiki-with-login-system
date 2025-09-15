package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var JWTSecret string
var JWTExpireHours int

func LoadJWTConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	JWTSecret = os.Getenv("JWT_SECRET")
	if JWTSecret == "" {
		log.Fatal("JWT_SECRET 未设置")
	}
	expireStr := os.Getenv("JWT_EXPIRE_HOURS")
	expire, err := strconv.Atoi(expireStr)
	if err != nil {
		log.Fatal("JWT_EXPIRE_HOURS 未设置或格式错误")
	}
	JWTExpireHours = expire
}
