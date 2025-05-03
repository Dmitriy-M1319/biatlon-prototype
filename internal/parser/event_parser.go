package parser

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Dmitriy-M1319/biatlon-prototype/internal/conveyor"
	"github.com/Dmitriy-M1319/biatlon-prototype/internal/models"
)

type EventParserImpl struct{}

func NewEventParserImpl() conveyor.EventParser {
	return &EventParserImpl{}
}

func (p *EventParserImpl) Parse(line string) (models.EventParsedDto, error) {
	parts := strings.Split(line, " ")
	result := models.EventParsedDto{}
	if len(parts) < 3 {
		return result, fmt.Errorf("Too few values in log line")
	}

	timestampStr := parts[0]
	timestamp, err := time.Parse("15:04:05.000", timestampStr[1:len(timestampStr)-1])
	if err != nil {
		return result, err
	}
	result.Timestamp = timestamp

	eventID, err := strconv.ParseUint(parts[1], 10, 32)
	if err != nil {
		return models.EventParsedDto{}, err
	}
	result.EventID = uint16(eventID)

	competitorID, err := strconv.ParseUint(parts[2], 10, 32)
	if err != nil {
		return models.EventParsedDto{}, err
	}
	result.CompetitorID = uint32(competitorID)

	switch eventID {
	case 2:
		result.ExtraParams = make(map[string]string)
		result.ExtraParams["startTime"] = parts[3]
		break
	case 5:
		result.ExtraParams = make(map[string]string)
		result.ExtraParams["firingRange"] = parts[3]
		break
	case 6:
		result.ExtraParams = make(map[string]string)
		result.ExtraParams["target"] = parts[3]
		break
	case 11:
		result.ExtraParams = make(map[string]string)
		result.ExtraParams["comment"] = strings.Join(parts[3:], " ")
		break
	}

	return result, nil
}
