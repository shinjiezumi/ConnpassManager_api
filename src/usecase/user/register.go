package user

import (
	"net/http"

	"gorm.io/gorm"

	cmerr "connpass-manager/common/error"
	"connpass-manager/common/general"
	"connpass-manager/db"
	"connpass-manager/domain/user"
)

// RegisterRequest ユーザー登録リクエスト
type RegisterRequest struct {
	Name     string `json:"name" validate:"required,max=255"`
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
	// 暗号化
	encryptedAddr := general.NewCryptString(req.Email)

	// 重複チェック
	repo := user.NewRepository(db.GetConnection())
	exists, err := repo.GetByEmail(encryptedAddr)
	if err != nil {
		return cmerr.NewApplicationError(http.StatusInternalServerError, "エラーが発生しました")
	} else if exists != nil {
		return cmerr.NewApplicationError(http.StatusConflict, "既に使用されているメールアドレスです")
	}

	// ユーザー登録する
	u := user.NewUser(req.Name, encryptedAddr, req.Password)
	if err := repo.Create(u); err != nil {
		return cmerr.NewApplicationError(http.StatusInternalServerError, "エラーが発生しました")
	}

	return nil
}
