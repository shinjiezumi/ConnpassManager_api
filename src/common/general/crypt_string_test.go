package general

import (
	"testing"

	"github.com/joho/godotenv"
)

const (
	text      = "hogehoge"
	encrypted = "pFvjxIozC-Y="
)

func TestNewCryptString(t *testing.T) {
	if err := godotenv.Load("/go/src/.env"); err != nil {
		panic("load env file failed")
	}

	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want CryptString
	}{
		{
			name: "正常系",
			args: args{
				text,
			},
			want: encrypted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCryptString(tt.args.str); got != tt.want {
				t.Errorf("NewCryptString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCryptString_Decrypt(t *testing.T) {
	tests := []struct {
		name string
		cs   CryptString
		want string
	}{
		{
			name: "正常系",
			cs:   CryptString(encrypted),
			want: text,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cs.Decrypt(); got != tt.want {
				t.Errorf("Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
