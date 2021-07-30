package user

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"connpass-manager/db"
	"connpass-manager/usecase/user"
)

// Login ログイン処理を行う
func Login(c echo.Context) error {
	req := new(user.LoginRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	// ユースケース実行
	if err := user.NewLoginUseCase(db.GetConnection()).Execute(c, req); err != nil {
		return err
	} else {
		return c.NoContent(http.StatusOK)
	}
}

// Logout ログアウト処理を行う
func Logout(c echo.Context) error {
	return c.String(http.StatusOK, "Logout")
}

// Register 会員登録を行う
func Register(c echo.Context) error {
	req := new(user.RegisterRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	// ユースケース実行
	if err := user.NewRegisterUseCase(db.GetConnection()).Execute(c, req); err != nil {
		return err
	} else {
		return c.NoContent(http.StatusOK)
	}
}

// PasswordReset パスワードリセットを行う
func PasswordReset(c echo.Context) error {
	return c.String(http.StatusOK, "PasswordReset")
}

// Withdraw 退会処理を行う
func Withdraw(c echo.Context) error {
	return c.String(http.StatusOK, "Withdraw")
}
