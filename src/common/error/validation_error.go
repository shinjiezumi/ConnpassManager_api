package error

import "encoding/json"

// ValidationError バリデーションエラー
type ValidationError struct {
	Code    int
	Type    Type
	Details []Detail
}

// Error .
func (v ValidationError) Error() string {
	j, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return string(j)
}

// Detail エラー詳細
type Detail struct {
	Field   string
	Message string
}

// NewValidationError バリデーションエラーを生成する
func NewValidationError(code int, details []Detail) *ValidationError {
	return &ValidationError{
		Code:    code,
		Type:    TypeValidation,
		Details: details,
	}
}
