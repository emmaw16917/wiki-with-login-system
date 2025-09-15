package database

import (
	"database/sql"
	"log"
	"wiki/backend/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	cfg := config.LoadDBConfig()
	dsn := cfg.DSN()
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatal("数据库无法访问:", err)
	}
}
