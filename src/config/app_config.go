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
