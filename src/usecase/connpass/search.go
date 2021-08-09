package connpass

import (
	"net/http"

	"gorm.io/gorm"

	cmerr "connpass-manager/common/error"
	"connpass-manager/domain/connpass"
)

// SearchRequest 検索リクエスト
type SearchRequest struct {
	Keyword string `json:"keyword" validate:"required"`
	Page    int    `json:"page" validate:"required"`
	Count   int    `json:"count" validate:"required,min=1,max=100"`
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
	events, err := connpass.NewSearcher().Search(req.Keyword, req.Page, req.Count)
	if err != nil {
		return nil, cmerr.NewApplicationError(http.StatusInternalServerError, "検索に失敗しました")
	}

	return &SearchResponse{
		Count:  len(events),
		Events: events,
	}, nil
}
