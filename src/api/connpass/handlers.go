package connpass

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"connpass-manager/db"
	"connpass-manager/usecase/connpass"
)

// Search イベントを検索する
func Search(c echo.Context) error {
	req := new(connpass.SearchRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	// ユースケース実行
	if res, err := connpass.NewSearchUseCase(db.GetConnection()).Execute(req); err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, res)
	}
}
