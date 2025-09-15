package main

import (
	"log"
	"wiki/backend/config"
	"wiki/backend/internal/api/router"
	"wiki/backend/pkg/database"
)

func main() {
	config.LoadJWTConfig()
	database.InitDB()
	r := router.SetupRouter()
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("服务启动失败:", err)
	}
}
