package user

import (
	"fmt"

	"gorm.io/gorm"
)

// RegisterRequest ユーザー登録リクエスト
type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,max=255"`
}

// RegisterUseCase ユーザー登録ユースケース
type RegisterUseCase struct {
	db *gorm.DB
}

// NewRegisterUseCase ユーザー登録ユースケースを生成する
func NewRegisterUseCase(db *gorm.DB) *RegisterUseCase {
	return &RegisterUseCase{
		db: db,
	}
}

// Execute ユーザー登録を実行する
func (uc *RegisterUseCase) Execute(req *RegisterRequest) error {
	// TODO 実装する
	fmt.Println(req)
	return nil
}
