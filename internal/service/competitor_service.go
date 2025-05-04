package service

import (
	"time"

	"github.com/Dmitriy-M1319/biatlon-prototype/internal/config"
	"github.com/Dmitriy-M1319/biatlon-prototype/internal/models"
)

type CompetitorService struct {
	CompetitorID      uint32
	currentMainLap    int
	currentPenaltyLap int
	startTime         time.Time
	MainLaps          []models.MainLap
	PenaltyLaps       []models.PenaltyLap
	StartFailed       bool
}

func RegisterCompetitor(cID uint32) *CompetitorService {
	return &CompetitorService{CompetitorID: cID,
		MainLaps: make([]models.MainLap, 0), PenaltyLaps: make([]models.PenaltyLap, 0)}
}

func (c *CompetitorService) PrepareResults() *models.CompetitorResult {
	r := &models.CompetitorResult{CompetitorID: c.CompetitorID}
	c.MainLaps = c.MainLaps[:len(c.MainLaps)-1]
	if (len(c.MainLaps) < config.Config().Laps) ||
		(c.MainLaps[len(c.MainLaps)-1].Speed == 0.0) {
		r.Status = "Not Finished"
	} else if c.StartFailed {
		r.Status = "Not Started"
	} else {
		r.DistanceTime = c.MainLaps[config.Config().Laps-1].End.Sub(c.startTime)
	}

	r.MainLaps = c.MainLaps

	for _, lap := range c.MainLaps {
		r.TotalHits += lap.HitsCount
		r.TotalShots += lap.ShotsCount
	}

	for _, pen := range c.PenaltyLaps {
		r.PenaltyTime += pen.Duration
	}
	r.PenaltyAvgSpeed = float32(config.Config().PenaltyLen*(r.TotalShots-r.TotalHits)) / float32(r.PenaltyTime.Seconds())

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
		c.processCompHit()
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
		c.MainLaps = append(c.MainLaps, models.MainLap{Start: e.Timestamp, ShotsCount: 5})
		c.currentMainLap += 1
	}
}

func (c *CompetitorService) processCompHit() {
	c.MainLaps[c.currentMainLap-1].HitsCount += 1
}

func (c *CompetitorService) processPenaltyLapEntered(e models.EventParsedDto) {
	c.PenaltyLaps = append(c.PenaltyLaps, models.PenaltyLap{Start: e.Timestamp})
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
	c.MainLaps[c.currentMainLap-1].Duration = dur
	c.MainLaps[c.currentMainLap-1].Speed = float32(config.Config().LapLen) / float32(dur.Seconds())

	c.MainLaps = append(c.MainLaps, models.MainLap{Start: e.Timestamp, ShotsCount: 5})
	c.currentMainLap += 1
}
