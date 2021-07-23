package request

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
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
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}
