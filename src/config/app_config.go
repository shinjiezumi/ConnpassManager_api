package config

import (
	"os"
)

type AppEnv string

const (
	AppEnvLocal      AppEnv = "local"
	AppEnvProduction AppEnv = "production"
)

// IsLocal ローカル環境かどうかを返す
func (e AppEnv) IsLocal() bool {
	return e == AppEnvLocal
}

// GetAppEnv 動作環境を返す
func GetAppEnv() AppEnv {
	env := os.Getenv("APP_ENV")
	if env == "" {
		panic("APP_ENVが設定されていません")
	}
	return AppEnv(env)
}

// GetAppURL アプリケーションのURLを返す
func GetAppURL() string {
	url := os.Getenv("APP_URL")
	if url == "" {
		panic("APP_URLが設定されていません")
	}
	return url
}
