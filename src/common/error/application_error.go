package error

import "encoding/json"

// ApplicationError アプリケーションエラー
type ApplicationError struct {
	Code    int    // ステータスコード
	Type    Type   // エラー種別
	Message string // エラーメッセージ
}

// Error .
func (a ApplicationError) Error() string {
	j, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}

	return string(j)
}

// NewApplicationError アプリケーションエラーを生成する
func NewApplicationError(code int, message string) *ApplicationError {
	return &ApplicationError{
		Code:    code,
		Type:    TypeApplication,
		Message: message,
	}
}
