package service

import (
	"time"

	"github.com/Dmitriy-M1319/biatlon-prototype/internal/models"
)

type CompetitorService struct {
	CompetitorID uint32
	Status       string
	DistanceTime time.Time
	HitsCount    int
	ShotsCount   int
}

func RegisterCompetitor(cID uint32) *CompetitorService {
	return &CompetitorService{CompetitorID: cID}
}

func (c *CompetitorService) ProcessEvent(e models.EventParsedDto) {
}
