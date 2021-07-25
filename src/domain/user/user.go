package user

import "time"

// User ユーザー
type User struct {
	ID              int       `gorm:"primaryKey"` // ユーザーID
	Name            *string   // ユーザー名
	Email           string    // メールアドレス
	EmailVerifiedAt *string   // メールアドレス認証日時
	Password        string    // パスワード
	RememberToken   *string   // パスワード再設定トークン
	CreatedAt       time.Time // 作成日時
	UpdatedAt       time.Time // 更新日時
}

// TableName テーブル名
func (u *User) TableName() string {
	return "users"
}

// NewUser ユーザーを生成する
func NewUser(email, password string) *User {
	return &User{
		Email:    email,
		Password: password,
	}
}
