package error

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// APIError APIエラー
type APIError struct {
	Type    Type      `json:"type"`
	Message *string   `json:"message,omitempty"`
	Details *[]Detail `json:"details,omitempty"`
}

// CustomHTTPErrorHandler エラーハンドラー
func CustomHTTPErrorHandler(err error, c echo.Context) {
	var t Type
	var code = http.StatusInternalServerError
	var message *string
	var details *[]Detail

	// エラー別に設定
	if ee, ok := err.(*ValidationError); ok {
		t = ee.Type
		code = ee.Code
		details = &ee.Details
	} else if ee, ok := err.(*ApplicationError); ok {
		t = ee.Type
		code = ee.Code
		message = &ee.Message
	}

	res := APIError{
		Type:    t,
		Message: message,
		Details: details,
	}

	if err := c.JSON(code, res); err != nil {
		return
	}
}
