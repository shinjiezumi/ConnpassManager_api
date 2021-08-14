package connpass

import (
	"net/http"

	"gorm.io/gorm"

	cmerr "connpass-manager/common/error"
	"connpass-manager/domain/connpass"
	"connpass-manager/domain/connpass/api"
)

// SearchRequest 検索リクエスト
type SearchRequest struct {
	Condition api.EventSearchQuery `json:"condition" validate:"required"`
	Page      int                  `json:"page" validate:"required"`
	Count     int                  `json:"count" validate:"required"`
}

// SearchResponse 検索レスポンス
type SearchResponse struct {
	Count  int               `json:"count"`
	Events []*connpass.Event `json:"event"`
}

// SearchUseCase イベント検索ユースケース
type SearchUseCase struct {
	db *gorm.DB
}

// NewSearchUseCase イベント検索ユースケースを生成する
func NewSearchUseCase(db *gorm.DB) *SearchUseCase {
	return &SearchUseCase{
		db: db,
	}
}

// Execute .
func (uc *SearchUseCase) Execute(req *SearchRequest) (*SearchResponse, error) {
	events, err := connpass.NewSearcher().Search(req.Condition, req.Page, req.Count)
	if err != nil {
		return nil, cmerr.NewApplicationError(http.StatusInternalServerError, "検索に失敗しました")
	}

	return &SearchResponse{
		Count:  len(events),
		Events: events,
	}, nil
}
