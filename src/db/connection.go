package db

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var connection *gorm.DB

// Initialize 初期化
func Initialize() {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	protocol := os.Getenv("DB_PROTOCOL")
	dbName := os.Getenv("DB_NAME")
	dsn := user + ":" + pass + "@" + protocol + "/" + dbName + "?charset=utf8mb4&parseTime=true&loc=Asia%2FTokyo"

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	connection = conn
}

// GetConnection DBコネクションを取得する
func GetConnection() *gorm.DB {
	return connection
}
