package connpass

import (
	"gorm.io/gorm"
)

// SearchRequest 検索リクエスト
type SearchRequest struct {
	Keyword string `json:"keyword" validate:"required"`
}

// SearchResponse 検索レスポンス
type SearchResponse struct {
	EventName string `json:"event_name"`
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
func (uc *SearchUseCase) Execute(req *SearchRequest) (SearchResponse, error) {
	// TODO 実装する
	return SearchResponse{
		EventName: "hogehoge",
	}, nil
}
