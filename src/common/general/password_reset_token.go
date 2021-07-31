package general

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"connpass-manager/common/date"
)

const (
	letters           = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	length            = 40
	separator         = "_"
	tokenExpiryMinute = 30 * time.Minute
)

// PasswordResetToken パスワード再設定トークン
type PasswordResetToken struct {
	token CryptString
}

// NewPasswordResetToken パスワード再設定トークンを生成する
func NewPasswordResetToken() PasswordResetToken {
	token := generateToken()
	encrypted := NewCryptString(token)
	return PasswordResetToken{
		token: encrypted,
	}
}

// IsExpired トークン期限が切れているかを返す
func (p PasswordResetToken) IsExpired() bool {
	decrypted := p.token.Decrypt()
	expiryDate, err := time.Parse(date.DefaultFormat, strings.Split(decrypted, separator)[1])
	if err != nil {
		panic(err)
	}
	return time.Now().Before(expiryDate)
}

func generateToken() string {
	randomStr := randomString(length)
	tokenExpiryDate := time.Now().Add(tokenExpiryMinute).Format(date.DefaultFormat)
	return fmt.Sprintf("%s%s%s", randomStr, separator, tokenExpiryDate)
}

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
