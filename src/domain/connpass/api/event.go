package api

import "time"

// @see https://connpass.com/about/api/

// EventSearchURL イベント検索ApiURL
const EventSearchURL = "https://connpass.com/api/v1/event/"

// EventSearchQuery イベント検索クエリ
type EventSearchQuery struct {
	EventID       *int    `json:"event_id"`       // イベントID
	Keyword       *string `json:"keyword"`        // キーワード(AND)
	KeywordOr     *string `json:"keyword_or"`     // キーワード(OR)
	Ym            *int    `json:"ym"`             // イベント開催年月
	Ymd           *int    `json:"ymd"`            // イベント開催年月日
	NickName      *string `json:"nickname"`       // 参加者のニックネーム
	OwnerNickname *string `json:"owner_nickname"` // 管理者のニックネーム
	SeriesID      *int    `json:"series_id"`      // グループID
	Start         int     `json:"-"`              // 検索の開始位置
	Order         *int    `json:"order"`          // 検索結果の表示順. 初期値:1、1:更新日時順、2:開催日時順、3:新着順
	Count         int     `json:"-"`              // 取得件数. 初期値:10、最小値:1、最大値:100
}

// EventResponse イベント検索レスポンス
type EventResponse struct {
	ResultsReturned  int     `json:"results_returned"`  // 含まれる検索結果の件数
	ResultsAvailable int     `json:"results_available"` // 検索結果の総件数
	ResultsStart     int     `json:"results_start"`     // 検索の開始位置
	Events           []Event `json:"events"`            // 検索結果のイベントリスト
}

// Event イベント
type Event struct {
	EventID          int       `json:"event_id"`           // イベントID
	Title            string    `json:"title"`              // タイトル
	Catch            string    `json:"catch"`              // キャッチ
	Description      string    `json:"description"`        // 概要(HTML形式)
	EventURL         string    `json:"event_url"`          // connpass.com上のURL
	HashTag          string    `json:"hash_tag"`           // twitterのハッシュタグ
	StartedAt        time.Time `json:"started_at"`         // イベント開催日時(ISO-8601形式)
	EndedAt          time.Time `json:"ended_at"`           // イベント終了日時(ISO-8601形式)
	Limit            int       `json:"limit"`              // 定員
	EventType        EventType `json:"event_type"`         // イベント参加タイプ
	Series           Series    `json:"series"`             // グループ
	Address          string    `json:"address"`            // 開催場所
	Place            string    `json:"place"`              // 開催会場
	Lat              string    `json:"lat"`                // 開催会場の緯度
	Lon              string    `json:"lon"`                // 開催会場の軽度
	OwnerID          int       `json:"owner_id"`           // 管理者のID
	OwnerNickname    string    `json:"owner_nickname"`     // 管理者のニックネーム
	OwnerDisplayName string    `json:"owner_display_name"` // 管理者の表示名
	Accepted         int       `json:"accepted"`           // 参加者数
	Waiting          int       `json:"waiting"`            // 補欠者数
	UpdateAt         time.Time `json:"updated_at"`         // 更新日時(ISO-8601形式)
}

// Series グループ
type Series struct {
	ID    int    `json:"id"`    // グループID
	Title string `json:"title"` // グループタイトル
	URL   string `json:"url"`   // グループのconnpass.com上のURL
}
