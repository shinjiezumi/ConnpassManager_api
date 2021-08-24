package connpass

import (
	"fmt"

	"github.com/labstack/echo/v4"

	cmerr "connpass-manager/common/error"
	"connpass-manager/common/request"
)

// APIカテゴリ
const apiCategory = "connpass"

// APIバージョン
const apiVersion = "v1"

// SetupRoutes イベント系APIのルーティングをセットアップする
func SetupRoutes(e *echo.Echo) {
	// バリデーター、エラーハンドラー設定
	e.Validator = request.NewValidator()
	e.HTTPErrorHandler = cmerr.CustomHTTPErrorHandler

	// ルーティング設定
	r := e.Group(fmt.Sprintf("/%s/%s", apiVersion, apiCategory))

	// connpassイベント検索
	r.GET("/search", Search)
}
