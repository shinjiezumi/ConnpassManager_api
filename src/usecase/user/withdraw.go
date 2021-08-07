package user

import (
	"net/http"

	"gorm.io/gorm"

	cmerr "connpass-manager/common/error"
	"connpass-manager/common/general"
	cmmail "connpass-manager/common/mail"
	"connpass-manager/db"
	"connpass-manager/domain/user"
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
func (uc *WithdrawUseCase) Execute(userID int) error {
	u, err := user.NewRepository(db.GetConnection()).GetByID(userID)
	if err != nil {
		return cmerr.NewApplicationError(http.StatusInternalServerError, "エラーが発生しました")
	} else if u == nil {
		return cmerr.NewApplicationError(http.StatusBadRequest, "ユーザー情報が取得できません")
	}

	tx := uc.db.Begin()

	email := u.Email
	// ユーザー削除
	if err := user.NewRepository(tx).Delete(u); err != nil {
		tx.Rollback()
		return cmerr.NewApplicationError(http.StatusInternalServerError, "エラーが発生しました")
	}

	// 退会完了メール送信
	if err := uc.sendWithdrawMail(email); err != nil {
		tx.Rollback()
		return cmerr.NewApplicationError(http.StatusInternalServerError, "エラーが発生しました")
	}

	if err := tx.Commit().Error; err != nil {
		return cmerr.NewApplicationError(http.StatusInternalServerError, "エラーが発生しました")
	}

	return nil
}

func (uc *WithdrawUseCase) sendWithdrawMail(email general.CryptString) error {
	toList := []string{email.Decrypt()}
	subject := "退会が完了しました"
	body := "退会処理が完了しました。\n ご利用ありがとうございました。\n"

	return cmmail.NewSender().SendTextMail(toList, subject, body)
}
