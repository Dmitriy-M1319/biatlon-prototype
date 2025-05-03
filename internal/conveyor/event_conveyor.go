package conveyor

import (
	"github.com/Dmitriy-M1319/biatlon-prototype/internal/models"
)

// Абстракции
// TODO: Интерфейс обработки логики соревнований Service

type EventParser interface {
	Parse(line string) (models.EventParsedDto, error)
}

type InputReader interface {
	ReadData() ([]string, error)
}

type OutputWriter interface {
	InsertLog(e models.EventParsedDto)
	InsertResult(line string)
	Write() error
}

type CompetitorService interface{}

type EventConveyor struct {
	Parser EventParser
	Writer OutputWriter
	Reader InputReader
}

func NewEventConveyor(p EventParser, w OutputWriter, r InputReader) *EventConveyor {
	return &EventConveyor{Parser: p, Writer: w, Reader: r}
}

func (c *EventConveyor) StartProcessEvents() error {
	events, err := c.Reader.ReadData()
	if err != nil {
		return nil
	}

	// Пока что просто вывод с подстановкой значений
	for _, event := range events {
		eventParsed, err := c.Parser.Parse(event)
		if err != nil {
			return err
		}

		c.Writer.InsertLog(eventParsed)
	}

	c.Writer.Write()
	return nil
}
