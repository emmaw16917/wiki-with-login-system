package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	User string
	Pass string
	Host string
	Port string
	Name string
}

func LoadDBConfig() *DBConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &DBConfig{
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Name: os.Getenv("DB_NAME"),
	}
}
func (cfg *DBConfig) DSN() string {
	return cfg.User + ":" + cfg.Pass + "@tcp(" + cfg.Host + ":" + cfg.Port + ")/" + cfg.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
}
