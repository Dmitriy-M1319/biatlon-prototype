package service

import (
	"strconv"
	"time"

	"github.com/Dmitriy-M1319/biatlon-prototype/internal/config"
	"github.com/Dmitriy-M1319/biatlon-prototype/internal/models"
)

// TODO: Makefile для сборки бинаря

type MainLap struct {
	Start      time.Time
	End        time.Time
	Speed      float32
	HitsCount  int
	ShotsCount int
}

type PenaltyLap struct {
	Start    time.Time
	End      time.Time
	Duration time.Duration
}

type CompetitorService struct {
	CompetitorID      uint32
	currentMainLap    int
	currentPenaltyLap int
	startTime         time.Time
	MainLaps          []MainLap
	PenaltyLaps       []PenaltyLap
	StartFailed       bool
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

func RegisterCompetitor(cID uint32) *CompetitorService {
	return &CompetitorService{CompetitorID: cID,
		MainLaps: make([]MainLap, 0), PenaltyLaps: make([]PenaltyLap, 0)}
}

func (c *CompetitorService) PrepareResults() *CompetitorResult {
	r := &CompetitorResult{MainLaps: c.MainLaps}
	if len(c.MainLaps) < config.Config().Laps {
		r.Status = "Not Finished"
	} else if c.StartFailed {
		r.Status = "Not Started"
	} else {
		r.DistanceTime = c.MainLaps[config.Config().Laps-1].End.Sub(c.startTime)
	}

	for _, lap := range c.MainLaps {
		r.TotalHits += lap.HitsCount
		r.TotalShots += lap.ShotsCount
	}

	for _, pen := range c.PenaltyLaps {
		r.PenaltyTime += pen.Duration
	}
	r.PenaltyAvgSpeed = float32(config.Config().PenaltyLen*len(c.PenaltyLaps)) / float32(r.PenaltyTime.Seconds())

	return r
}

func (c *CompetitorService) ProcessEvent(e models.EventParsedDto) {
	switch e.EventID {
	case 2:
		c.processSetStartTime(e)
		break
	case 4:
		c.processCompHasStarted(e)
		break
	case 6:
		c.processCompHit(e)
		break
	case 8:
		c.processPenaltyLapEntered(e)
		break
	case 9:
		c.processPenaltyLapLeft(e)
		break
	case 10:
		c.processMainLapEnd(e)
		break
	}
}

func (c *CompetitorService) processSetStartTime(e models.EventParsedDto) {
	t, _ := time.Parse("15:04:05.000", e.ExtraParams["startTime"])
	c.startTime = t
}

func (c *CompetitorService) processCompHasStarted(e models.EventParsedDto) {
	if e.Timestamp.Sub(c.startTime) > config.Config().StartDeltaTime {
		c.StartFailed = true
	} else {
		c.MainLaps = append(c.MainLaps, MainLap{Start: e.Timestamp})
		c.currentMainLap += 1
	}
}

func (c *CompetitorService) processCompHit(e models.EventParsedDto) {
	c.MainLaps[c.currentMainLap-1].HitsCount += 1
	targetNumber, _ := strconv.Atoi(e.ExtraParams["target"])
	c.MainLaps[c.currentMainLap-1].ShotsCount =
		max(c.MainLaps[c.currentMainLap-1].ShotsCount, targetNumber)
}

func (c *CompetitorService) processPenaltyLapEntered(e models.EventParsedDto) {
	c.PenaltyLaps = append(c.PenaltyLaps, PenaltyLap{Start: e.Timestamp})
	c.currentPenaltyLap += 1
}

func (c *CompetitorService) processPenaltyLapLeft(e models.EventParsedDto) {
	c.PenaltyLaps[c.currentPenaltyLap-1].End = e.Timestamp
	c.PenaltyLaps[c.currentPenaltyLap-1].Duration =
		e.Timestamp.Sub(c.PenaltyLaps[c.currentPenaltyLap-1].Start)
}

func (c *CompetitorService) processMainLapEnd(e models.EventParsedDto) {
	c.MainLaps[c.currentMainLap-1].End = e.Timestamp
	dur := e.Timestamp.Sub(c.MainLaps[c.currentMainLap-1].Start)
	c.MainLaps[c.currentMainLap-1].Speed = float32(config.Config().LapLen) / float32(dur.Seconds())
}
