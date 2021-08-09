package api

// EventType イベント種別
type EventType string

const (
	EventTypeParticipation EventType = "participation" // connpassで参加受付あり
	EventTypeAdvertisement EventType = "advertisement" // 告知のみ
)
