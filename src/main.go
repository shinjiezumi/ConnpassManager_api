package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"connpass-manager/api/connpass"
	"connpass-manager/api/user"
	"connpass-manager/common/session"
	"connpass-manager/config"
	"connpass-manager/db"
	"connpass-manager/logger"
)

// APIポート
const apiPort = 1323

func main() {
	e := echo.New()

	initialize(e)

	// ROOT
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "connpass-manager")
	})

	// ルーティングの設定
	user.SetupRoutes(e)
	connpass.SetupRoutes(e)

	log.Println(fmt.Sprintf("start api at %s env", config.GetAppEnv()))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", apiPort)))
}

func initialize(e *echo.Echo) {
	if err := godotenv.Load(".env"); err != nil {
		panic("load env file failed")
	}

	// アクセスログの設定
	logger.Setup(e)
	e.Use(middleware.Recover())
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookieHTTPOnly: true,
	}))

	// データベース初期化
	db.Initialize()

	// セッション初期化
	session.Initialize(e)
}
