package connpass

import (
	"time"

	"connpass-manager/domain/connpass/api"
)

// Event connpassイベント
type Event struct {
	Title            string     `json:"title"`              // タイトル
	Description      string     `json:"description"`        // 概要
	URL              string     `json:"url"`                // URL
	StartedAt        time.Time  `json:"started_at"`         // イベント開催日時
	EndedAt          time.Time  `json:"ended_at"`           // イベント終了日時
	Limit            int        `json:"limit"`              // 定員
	Series           api.Series `json:"series"`             // グループ
	Address          string     `json:"address"`            // 開催場所
	Place            string     `json:"place"`              // 開催会場
	OwnerDisplayName string     `json:"owner_display_name"` // 管理者表示名
	Accepted         int        `json:"accepted"`           // 参加者数
	Waiting          int        `json:"waiting"`            // 補欠者数
}
