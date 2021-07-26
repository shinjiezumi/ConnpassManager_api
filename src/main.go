package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"connpass-manager/api/user"
	"connpass-manager/config"
	"connpass-manager/db"
	"connpass-manager/logger"
)

// APIポート
const apiPort = 1323

func main() {
	e := echo.New()

	if err := godotenv.Load(".env"); err != nil {
		panic("load env file failed")
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

	user.SetupRoutes(e)

	log.Println(fmt.Sprintf("start api at %s env", config.GetAppEnv()))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", apiPort)))
}
