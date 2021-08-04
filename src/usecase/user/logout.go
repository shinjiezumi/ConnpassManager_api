package user

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"connpass-manager/common/session"
)

// LogoutUseCase ユーザーログアウトユースケース
type LogoutUseCase struct {
	db *gorm.DB
}

// NewLogoutUseCase ユーザーログアウトユースケースを生成する
func NewLogoutUseCase(db *gorm.DB) *LogoutUseCase {
	return &LogoutUseCase{
		db: db,
	}
}

// Execute ユーザーログアウトを実行する
func (uc *LogoutUseCase) Execute(c echo.Context) error {
	session.Destroy(c)
	return nil
}
