package user

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"

	cmerr "connpass-manager/common/error"
	"connpass-manager/common/general"
	cmmail "connpass-manager/common/mail"
	"connpass-manager/config"
	"connpass-manager/db"
	"connpass-manager/domain/user"
	"connpass-manager/domain/vo"
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

	// トークン生成＋DB保存
	token := vo.NewPasswordResetToken(req.Email)
	u.SetPasswordResetToken(token)
	tx := db.GetConnection().Begin()
	if err := user.NewRepository(tx).Save(u); err != nil {
		tx.Rollback()
		return cmerr.NewApplicationError(http.StatusInternalServerError, "エラーが発生しました")
	}

	// メール送信
	if err := uc.sendPasswordResetMail(u.Email, token); err != nil {
		tx.Rollback()
		panic(err)
	}

	if err := tx.Commit().Error; err != nil {
		return cmerr.NewApplicationError(http.StatusInternalServerError, "エラーが発生しました")
	}

	return nil
}

func (uc *PasswordResetUseCase) sendPasswordResetMail(email general.CryptString, token vo.PasswordResetToken) error {
	toList := []string{email.Decrypt()}
	subject := "パスワード再設定"
	template := "パスワードの再設定は以下URLページより、%d分以内に行ってください。\n %s"
	passwordResetURL := fmt.Sprintf("%s/password_reset?token=%s", config.GetAppURL(), token)
	body := fmt.Sprintf(template, vo.TokenExpiryMinute, passwordResetURL)

	return cmmail.NewSender().SendTextMail(toList, subject, body)
}
