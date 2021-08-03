package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	cmerr "connpass-manager/common/error"
	"connpass-manager/common/general"
	"connpass-manager/common/session"
	"connpass-manager/db"
	"connpass-manager/domain/user"
	"connpass-manager/domain/vo"
)

// PasswordResetRequest パスワード再設定リクエスト
type PasswordResetRequest struct {
	Token    string `json:"token" validate:"required"`
	Password string `json:"password" validate:"required,max=255"`
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
func (uc *PasswordResetUseCase) Execute(c echo.Context, req *PasswordResetRequest) error {
	token := vo.PasswordResetToken(req.Token)
	// 有効期限チェック
	if token.IsExpired() {
		return cmerr.NewApplicationError(http.StatusUnauthorized, "有効期限が切れています。初めからやり直してください")
	}

	repo := user.NewRepository(db.GetConnection())
	u, err := repo.GetByEmail(token.GetEmail())
	if err != nil {
		return cmerr.NewApplicationError(http.StatusInternalServerError, "エラーが発生しました")
	} else if u == nil {
		// FIXME 再設定後に退会など。汎用的にするか？
		return cmerr.NewApplicationError(http.StatusUnauthorized, "メールアドレスが誤っています")
	} else if u.RememberToken == nil || *u.RememberToken != token {
		return cmerr.NewApplicationError(http.StatusUnauthorized, "認証できませんでした。初めからやり直してください")
	}

	// 新しいパスワードを保存
	hashPassword := general.NewHashString(req.Password)
	u.ResetPassword(hashPassword)
	if err := repo.Save(u); err != nil {
		return cmerr.NewApplicationError(http.StatusInternalServerError, "エラーが発生しました")
	}

	session.SaveUserID(c, u.ID)

	return nil
}
