package user

import (
	"errors"

	"gorm.io/gorm"

	"connpass-manager/common/general"
)

// Repository ユーザーリポジトリ
type Repository struct {
	db *gorm.DB
}

// NewRepository ユーザーリポジトリを生成する
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// GetByEmail メールアドレスで取得する
func (r *Repository) GetByEmail(email general.CryptString) (*User, error) {
	var ret User

	if err := r.db.Where("email = ?", email).First(&ret).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &ret, nil
}

// GetByEmailAndPassword メールアドレスとパスワードで取得する
func (r *Repository) GetByEmailAndPassword(email general.CryptString, password general.HashString) (*User, error) {
	var ret User

	if err := r.db.
		Where("email = ?", email).
		Where("password = ?", password).
		First(&ret).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &ret, nil
}

// Create ユーザーを登録する
func (r *Repository) Create(u *User) error {
	if err := r.db.Create(u).Error; err != nil {
		return err
	}

	return nil
}
