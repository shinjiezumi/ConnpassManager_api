package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"connpass-manager/config"
	"connpass-manager/db"
	"connpass-manager/logger"
)

// APIポート
const apiPort = 1323

func main() {
	e := echo.New()

	if os.Getenv("APP_ENV") == "local" {
		if err := godotenv.Load(".env"); err != nil {
			panic("load env file failed")
		}
	}

	// アクセスログの設定
	logger.Setup(e)
	e.Use(middleware.Recover())

	// データベースセットアップ
	db.Initialize()

	// ROOT
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "connpass-manager")
	})

	log.Println(fmt.Sprintf("start api at %s env", config.GetAppEnv()))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", apiPort)))
}
