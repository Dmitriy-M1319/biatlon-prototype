package models

import "time"

type EventParsedDto struct {
	Timestamp    time.Time
	CompetitorID uint32
	EventID      uint16
	ExtraParams  map[string]string
}
