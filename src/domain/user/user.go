package user

import "time"

// User ユーザー
type User struct {
	ID              int `gorm:"primaryKey"`
	Name            string
	Email           string
	EmailVerifiedAt *string
	Password        string
	RememberToken   *string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// TableName テーブル名
func (u *User) TableName() string {
	return "users"
}
