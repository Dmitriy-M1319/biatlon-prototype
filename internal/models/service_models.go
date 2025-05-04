package models

import "time"

type MainLap struct {
	Start      time.Time
	End        time.Time
	Duration   time.Duration
	Speed      float32
	HitsCount  int
	ShotsCount int
}

type PenaltyLap struct {
	Start    time.Time
	End      time.Time
	Duration time.Duration
}

type CompetitorResult struct {
	CompetitorID    uint32
	Status          string
	DistanceTime    time.Duration
	MainLaps        []MainLap
	PenaltyTime     time.Duration
	PenaltyAvgSpeed float32
	TotalHits       int
	TotalShots      int
}

type ByDistanceTime []*CompetitorResult

func (a ByDistanceTime) Len() int {
	return len(a)
}

func (a ByDistanceTime) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByDistanceTime) Less(i, j int) bool {
	return a[i].DistanceTime < a[j].DistanceTime
}
