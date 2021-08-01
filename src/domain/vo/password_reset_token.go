package vo

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"connpass-manager/common/date"
	"connpass-manager/common/general"
)

const (
	letters           = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	length            = 40
	separator         = "_"
	TokenExpiryMinute = 30
)

// PasswordResetToken パスワード再設定トークン
type PasswordResetToken general.CryptString

// NewPasswordResetToken パスワード再設定トークンを生成する
func NewPasswordResetToken() PasswordResetToken {
	token := generateToken()
	encrypted := general.NewCryptString(token)
	return PasswordResetToken(encrypted)
}

// IsExpired トークン期限が切れているかを返す
func (p *PasswordResetToken) IsExpired() bool {
	decrypted := general.CryptString(*p).Decrypt()
	expiryDate, err := time.Parse(date.DefaultFormat, strings.Split(decrypted, separator)[1])
	if err != nil {
		panic(err)
	}
	return time.Now().Before(expiryDate)
}

func generateToken() string {
	randomStr := randomString(length)
	tokenExpiryDate := time.Now().Add(TokenExpiryMinute * time.Minute).Format(date.DefaultFormat)
	return fmt.Sprintf("%s%s%s", randomStr, separator, tokenExpiryDate)
}

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
