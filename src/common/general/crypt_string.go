package general

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"os"
)

// CryptString 暗号化文字列
type CryptString string

// NewCryptString 暗号化文字列を生成する
func NewCryptString(str string) CryptString {
	c := newCipher()
	iv := loadIV()

	b := []byte(str)
	cfb := cipher.NewCFBEncrypter(c, []byte(iv))
	encrypted := make([]byte, len(b))
	cfb.XORKeyStream(encrypted, b)
	encoded := base64.URLEncoding.EncodeToString(encrypted)

	return CryptString(encoded)
}

// Decrypt 復号化する
func (cs CryptString) Decrypt() string {
	c := newCipher()
	iv := loadIV()

	decoded, err := base64.URLEncoding.DecodeString(string(cs))
	if err != nil {
		panic(err)
	}

	cfbDec := cipher.NewCFBDecrypter(c, []byte(iv))
	decrypted := make([]byte, len(decoded))
	cfbDec.XORKeyStream(decrypted, decoded)

	return string(decrypted)
}

func newCipher() cipher.Block {
	cryptKey := os.Getenv("CRYPT_KEY")
	if cryptKey == "" {
		panic("CRYPT_KEY is empty")
	}

	c, err := aes.NewCipher([]byte(cryptKey))
	if err != nil {
		panic(err)
	}

	return c
}

func loadIV() string {
	iv := os.Getenv("CRYPT_IV")
	if iv == "" {
		panic("CRYPT_IV is empty")
	}
	return iv
}
