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
	length            = 10
	separator         = "_"
	TokenExpiryMinute = 30
)

// PasswordResetToken パスワード再設定トークン
type PasswordResetToken general.CryptString

// NewPasswordResetToken パスワード再設定トークンを生成する
func NewPasswordResetToken(email string) PasswordResetToken {
	token := generateToken(email)
	encrypted := general.NewCryptString(token)
	return PasswordResetToken(encrypted)
}

// GetEmail トークンからメールアドレスを返す
func (p *PasswordResetToken) GetEmail() general.CryptString {
	decrypted := general.CryptString(*p).Decrypt()
	email := strings.Split(decrypted, separator)[1]

	return general.NewCryptString(email)
}

// IsExpired トークン期限が切れているかを返す
func (p *PasswordResetToken) IsExpired() bool {
	decrypted := general.CryptString(*p).Decrypt()
	expiryDate, err := time.Parse(date.DefaultFormat, strings.Split(decrypted, separator)[2])
	if err != nil {
		panic(err)
	}
	return time.Now().Before(expiryDate)
}

func generateToken(email string) string {
	randomStr := randomString(length)
	tokenExpiryDate := time.Now().Add(TokenExpiryMinute * time.Minute).Format(date.DefaultFormat)
	return fmt.Sprintf("%s", strings.Join([]string{randomStr, email, tokenExpiryDate}, separator))
}

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
