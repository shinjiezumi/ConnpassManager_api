package user

import (
	"fmt"

	"github.com/labstack/echo/v4"

	cmerr "connpass-manager/common/error"
	"connpass-manager/common/request"
)

const apiCategory = "user"
const apiVersion = "v1"

// SetupRoutes ユーザー系APIのルーティングをセットアップする
func SetupRoutes(e *echo.Echo) {
	// バリデーター、エラーハンドラー設定
	e.Validator = request.NewValidator()
	e.HTTPErrorHandler = cmerr.CustomHTTPErrorHandler

	// ルーティング設定
	r := e.Group(fmt.Sprintf("/%s/%s", apiVersion, apiCategory))

	r.POST("/login", Login)
	r.POST("/logout", Logout)
	r.POST("/register", Register)
	r.POST("/forgot_password", ForgotPassword)
	r.POST("/password_reset", PasswordReset)
	r.POST("/withdraw", Withdraw)
}
