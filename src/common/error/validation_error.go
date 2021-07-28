package error

import "encoding/json"

// ValidationError バリデーションエラー
type ValidationError struct {
	Code    int      // ステータスコード
	Type    Type     // エラー種別
	Details []Detail // エラー詳細
}

// Error エラーメッセージを返す
func (v ValidationError) Error() string {
	j, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return string(j)
}

// Detail エラー詳細
type Detail struct {
	Field   string // エラー項目
	Message string // エラーメッセージ。現状タグのみ
}

// NewValidationError バリデーションエラーを生成する
func NewValidationError(code int, details []Detail) *ValidationError {
	return &ValidationError{
		Code:    code,
		Type:    TypeValidation,
		Details: details,
	}
}
