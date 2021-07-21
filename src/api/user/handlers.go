package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Login ログイン処理を行います
func Login(c echo.Context) error {
	return c.String(http.StatusOK, "Login")
}

// Logout ログアウト処理を行います
func Logout(c echo.Context) error {
	return c.String(http.StatusOK, "Logout")
}

// Register 会員登録を行います
func Register(c echo.Context) error {
	return c.String(http.StatusOK, "Register")
}

// PasswordReset パスワードリセットを行います
func PasswordReset(c echo.Context) error {
	return c.String(http.StatusOK, "PasswordReset")
}
