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

// ForgotPasswordRequest パスワード再設定要求リクエスト
type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email,max=255"`
}

// ForgotPasswordUseCase パスワード再設定要求ユースケース
type ForgotPasswordUseCase struct {
	db *gorm.DB
}

// NewForgotPasswordUseCase パスワード再設定要求ユースケースを生成する
func NewForgotPasswordUseCase(db *gorm.DB) *ForgotPasswordUseCase {
	return &ForgotPasswordUseCase{
		db: db,
	}
}

// Execute パスワード再設定要求を実行する
func (uc *ForgotPasswordUseCase) Execute(req *ForgotPasswordRequest) error {
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

func (uc *ForgotPasswordUseCase) sendPasswordResetMail(email general.CryptString, token vo.PasswordResetToken) error {
	toList := []string{email.Decrypt()}
	subject := "パスワード再設定"
	template := "パスワードの再設定は以下URLページより、%d分以内に行ってください。\n %s"
	passwordResetURL := fmt.Sprintf("%s/password_reset?token=%s", config.GetAppURL(), token)
	body := fmt.Sprintf(template, vo.TokenExpiryMinute, passwordResetURL)

	return cmmail.NewSender().SendTextMail(toList, subject, body)
}
