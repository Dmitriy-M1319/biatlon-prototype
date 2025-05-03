package models

import "time"

type EventParsedDto struct {
	Timestamp    time.Time
	CompetitorID uint32
	EventID      uint16
	ExtraParams  map[string]string
}

type Config struct {
	Laps        int    `json:"laps"`
	LapLen      int    `json:"lapLen"`
	PenaltyLen  int    `json:"penaltyLen"`
	FiringLines int    `json:"firingLines"`
	Start       string `json:"start"`
	StartDelta  string `json:"startDelta"`
}
