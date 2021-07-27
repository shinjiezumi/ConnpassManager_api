package general

import (
	"encoding/base64"
	"os"

	"golang.org/x/crypto/scrypt"
)

// HashString ハッシュ文字列
type HashString string

// NewHashString ハッシュ文字列を生成する
func NewHashString(str string) HashString {
	salt := loadSalt()

	hash, err := scrypt.Key([]byte(str), []byte(salt), 16384, 8, 1, 32)
	if err != nil {
		panic(err)
	}

	h := base64.URLEncoding.EncodeToString(hash[:])
	return HashString(h)
}

func loadSalt() string {
	salt := os.Getenv("HASH_SALT")
	if salt == "" {
		panic("HASH_SALT is empty")
	}

	return salt
}
