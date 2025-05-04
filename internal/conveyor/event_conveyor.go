package conveyor

import (
	"sort"

	"github.com/Dmitriy-M1319/biatlon-prototype/internal/models"
	"github.com/Dmitriy-M1319/biatlon-prototype/internal/service"
)

type EventParser interface {
	Parse(line string) (models.EventParsedDto, error)
}

type EventsInputReader interface {
	ReadData() ([]string, error)
}

type OutputWriter interface {
	InsertLog(e models.EventParsedDto)
	InsertResult(r *models.CompetitorResult)
	Write() error
}

type EventConveyor struct {
	Parser      EventParser
	Writer      OutputWriter
	Reader      EventsInputReader
	Competitors map[uint32]*service.CompetitorService
}

func NewEventConveyor(p EventParser,
	w OutputWriter, r EventsInputReader) *EventConveyor {
	return &EventConveyor{Parser: p, Writer: w, Reader: r,
		Competitors: make(map[uint32]*service.CompetitorService)}
}

func (c *EventConveyor) StartProcessEvents() error {
	events, err := c.Reader.ReadData()
	if err != nil {
		return nil
	}

	for _, event := range events {
		eventParsed, err := c.Parser.Parse(event)
		if err != nil {
			return err
		}

		if eventParsed.EventID == 1 {
			c.Competitors[eventParsed.CompetitorID] = service.RegisterCompetitor(eventParsed.CompetitorID)
		} else {
			c.Competitors[eventParsed.CompetitorID].ProcessEvent(eventParsed)
		}

		c.Writer.InsertLog(eventParsed)
	}

	sortedResults := make([]*models.CompetitorResult, 0)
	for _, serv := range c.Competitors {
		sortedResults = append(sortedResults, serv.PrepareResults())
	}

	sort.Sort(models.ByDistanceTime(sortedResults))
	for _, res := range sortedResults {
		c.Writer.InsertResult(res)
	}

	c.Writer.Write()
	return nil
}
