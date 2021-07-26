package general

import (
	"crypto/sha256"
	"encoding/base64"
)

// HashString ハッシュ文字列
type HashString string

// NewHashString ハッシュ文字列を生成する
func NewHashString(str string) HashString {
	b := sha256.Sum256([]byte(str))
	h := base64.URLEncoding.EncodeToString(b[:])
	return HashString(h)
}
