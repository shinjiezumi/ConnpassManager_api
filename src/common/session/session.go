package session

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Initialize(e *echo.Echo) {
	sessionKey := os.Getenv("SESSION_KEY")
	if sessionKey == "" {
		panic("SESSION_KEY is empty")
	}
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(sessionKey))))
}
