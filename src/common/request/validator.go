package request

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"

	cmerr "connpass-manager/common/error"
)

// CustomValidator .
type CustomValidator struct {
	validator *validator.Validate
}

// NewValidator .
func NewValidator() echo.Validator {
	return &CustomValidator{
		validator: validator.New(),
	}
}

// Validate .
func (v CustomValidator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		details := make([]cmerr.Detail, 0, len(err.(validator.ValidationErrors)))
		for _, err := range err.(validator.ValidationErrors) {
			details = append(details, cmerr.Detail{
				// FIXME 必要に応じて丁寧に出す
				Field:   err.Field(),
				Message: err.Tag(),
			})
		}
		return cmerr.NewValidationError(http.StatusBadRequest, details)
	}

	return nil
}
