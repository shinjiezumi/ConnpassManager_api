package user

import (
	"time"

	"connpass-manager/common/general"
	"connpass-manager/domain/vo"
)

// User ユーザー
type User struct {
	ID              int                    `gorm:"primaryKey"` // ユーザーID
	Name            string                 // ユーザー名
	Email           general.CryptString    // メールアドレス(暗号化)
	EmailVerifiedAt *string                // メールアドレス認証日時
	Password        general.HashString     // パスワード
	RememberToken   *vo.PasswordResetToken // パスワード再設定トークン
	CreatedAt       time.Time              // 作成日時
	UpdatedAt       time.Time              // 更新日時
}

// TableName テーブル名を返す
func (u *User) TableName() string {
	return "users"
}

// NewUser ユーザーを生成する
func NewUser(name string, email general.CryptString, password general.HashString) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}

// SetPasswordResetToken パスワードリセットトークンを設定する
func (u *User) SetPasswordResetToken(token vo.PasswordResetToken) {
	u.RememberToken = &token
}

// ResetPassword パスワードをリセットする
func (u *User) ResetPassword(newPassword general.HashString) {
	u.Password = newPassword
	u.RememberToken = nil
}
