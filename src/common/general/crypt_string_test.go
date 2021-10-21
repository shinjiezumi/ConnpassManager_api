package general

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

const (
	text      = "hogehoge"
	encrypted = "pFvjxIozC-Y="
)

func init() {
	appRoot := os.Getenv("CM_APP_ROOT")
	if appRoot == "" {
		panic("CM_APP_ROOT is empty")
	}

	if err := godotenv.Load(fmt.Sprintf("%s/.env", appRoot)); err != nil {
		panic("load env file failed")
	}
}

func TestNewCryptString(t *testing.T) {
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
