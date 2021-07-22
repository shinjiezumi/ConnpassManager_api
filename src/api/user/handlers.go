package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Login ログイン処理を行う
func Login(c echo.Context) error {
	return c.String(http.StatusOK, "Login")
}

// Logout ログアウト処理を行う
func Logout(c echo.Context) error {
	return c.String(http.StatusOK, "Logout")
}

// Register 会員登録を行う
func Register(c echo.Context) error {
	return c.String(http.StatusOK, "Register")
}

// PasswordReset パスワードリセットを行う
func PasswordReset(c echo.Context) error {
	return c.String(http.StatusOK, "PasswordReset")
}

// Withdraw 退会処理を行う
func Withdraw(c echo.Context) error {
	return c.String(http.StatusOK, "Withdraw")
}
