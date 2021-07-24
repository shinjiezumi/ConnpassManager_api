package error

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// APIError .
type APIError struct {
	Type    Type      `json:"type"`
	Message *string   `json:"message,omitempty"`
	Details *[]Detail `json:"details,omitempty"`
}

// CustomHTTPErrorHandler .
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
