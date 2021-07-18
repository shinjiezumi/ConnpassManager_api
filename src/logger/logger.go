package logger

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const logFilePath = "/var/log/webapps/connpass-manager/log.txt"

// Setup ロガーをセットアップします
func Setup(e *echo.Echo) {
	fp, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "${time_custom} method=${method}, uri=${uri}, status=${status}, error=${error}, user_agent=${user_agent}, remote_ip=${remote_ip}\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
		Output:           fp,
	}))
}
