package user

import (
	"net/http"

	"gorm.io/gorm"

	cmerr "connpass-manager/common/error"
	"connpass-manager/common/general"
	"connpass-manager/db"
	"connpass-manager/domain/user"
)

// PasswordResetRequest パスワード再設定リクエスト
type PasswordResetRequest struct {
	Email string `json:"email" validate:"required,email,max=255"`
}

// PasswordResetUseCase パスワード再設定ユースケース
type PasswordResetUseCase struct {
	db *gorm.DB
}

// NewPasswordResetUseCase パスワード再設定ユースケースを生成する
func NewPasswordResetUseCase(db *gorm.DB) *PasswordResetUseCase {
	return &PasswordResetUseCase{
		db: db,
	}
}

// Execute パスワード再設定を実行する
func (uc *PasswordResetUseCase) Execute(req *PasswordResetRequest) error {
	// 暗号化
	encryptedAddr := general.NewCryptString(req.Email)

	// ユーザーチェックする
	repo := user.NewRepository(db.GetConnection())
	u, err := repo.GetByEmail(encryptedAddr)
	if err != nil {
		return cmerr.NewApplicationError(http.StatusInternalServerError, "エラーが発生しました")
	} else if u == nil {
		return cmerr.NewApplicationError(http.StatusBadRequest, "メールアドレスが誤っています")
	}

	token := general.NewPasswordResetToken()
	u.SetPasswordResetToken(token)

	tx := db.GetConnection().Begin()
	if err := user.NewRepository(tx).Save(u); err != nil {
		tx.Rollback()
		return cmerr.NewApplicationError(http.StatusInternalServerError, "エラーが発生しました")
	}

	// TODO メール送信

	if err := tx.Commit(); err != nil {
		return cmerr.NewApplicationError(http.StatusInternalServerError, "エラーが発生しました")
	}

	return nil
}
