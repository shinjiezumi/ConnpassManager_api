package user

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// WithdrawUseCase ユーザー退会ユースケース
type WithdrawUseCase struct {
	db *gorm.DB
}

// NewWithdrawUseCase ユーザー退会ユースケースを生成する
func NewWithdrawUseCase(db *gorm.DB) *WithdrawUseCase {
	return &WithdrawUseCase{
		db: db,
	}
}

// Execute ユーザー退会を実行する
func (uc *WithdrawUseCase) Execute(c echo.Context) error {
	// TODO: 実装する
	// 退会処理

	// 退会完了メール送信

	return nil
}
