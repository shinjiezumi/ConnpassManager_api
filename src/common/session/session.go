package session

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const (
	sessionName       = "cm_user"
	sessionKeyUserID  = "user_id"
	sessionExpiryDate = 180 * 86400 // 180日
)

// Initialize 初期化する
func Initialize(e *echo.Echo) {
	sessionKey := os.Getenv("SESSION_SECRET")
	if sessionKey == "" {
		panic("SESSION_SECRET is empty")
	}
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(sessionKey))))
}

// GetUserID ユーザーIDを取得する
func GetUserID(c echo.Context) *int {
	sess, err := session.Get(sessionName, c)
	if err != nil {
		panic(err)
	}

	if userID, ok := sess.Values[sessionKeyUserID]; !ok {
		return nil
	} else {
		ret := userID.(int)
		return &ret
	}
}

// SaveUserID ユーザーIDを保存する
func SaveUserID(c echo.Context, userID int) {
	sess, err := session.Get(sessionName, c)
	if err != nil {
		panic(err)
	}
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   sessionExpiryDate,
		HttpOnly: true,
	}
	sess.Values[sessionKeyUserID] = userID

	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		panic(err)
	}
}

// Destroy セッションを破棄する
func Destroy(c echo.Context) {
	sess, err := session.Get(sessionName, c)
	if err != nil {
		panic(err)
	}
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	sess.Values[sessionKeyUserID] = ""

	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		panic(err)
	}
}
