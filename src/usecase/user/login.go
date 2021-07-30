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
)

// LoginRequest ユーザーログインリクエスト
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,max=255"`
}

// LoginUseCase ユーザーログインユースケース
type LoginUseCase struct {
	db *gorm.DB
}

// NewLoginUseCase ユーザーログインユースケースを生成する
func NewLoginUseCase(db *gorm.DB) *LoginUseCase {
	return &LoginUseCase{
		db: db,
	}
}

// Execute ユーザーログインを実行する
func (uc *LoginUseCase) Execute(c echo.Context, req *LoginRequest) error {
	// 暗号化
	encryptedAddr := general.NewCryptString(req.Email)
	// ハッシュ化
	hashPassword := general.NewHashString(req.Password)

	// ユーザーログインする
	repo := user.NewRepository(db.GetConnection())
	u, err := repo.GetByEmailAndPassword(encryptedAddr, hashPassword)
	if err != nil {
		return cmerr.NewApplicationError(http.StatusInternalServerError, "エラーが発生しました")
	} else if u == nil {
		return cmerr.NewApplicationError(http.StatusUnauthorized, "メールアドレス、またはパスワードが誤っています")
	}

	session.SaveUserID(c, u.ID)

	return nil
}
