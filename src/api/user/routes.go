package user

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

const apiCategory = "user"
const apiVersion = "v1"

// SetupRoutes ユーザー系APIのルーティングをセットアップする
func SetupRoutes(e *echo.Echo) {
	r := e.Group(fmt.Sprintf("/%s/%s", apiVersion, apiCategory))

	r.POST("/login", Login)
	r.POST("/logout", Logout)
	r.POST("/register", Register)
	r.POST("/password_reset", PasswordReset)
	r.POST("/withdraw", Withdraw)
}
