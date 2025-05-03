package conveyor

import (
	"fmt"

	"github.com/Dmitriy-M1319/biatlon-prototype/internal/models"
)

// Абстракции
// TODO: Отдельная подсистема логирования (OutputWriter, 2 метода для обычных логов и для результатов)
// TODO: Сюда же отдельная подсистема ввода (InputReader, желательно все объединить в какой нибудь пакет io
// TODO: Интерфейс обработки логики соревнований Service

type OutputWriter interface {
	InsertLog(line string)
	InsertResult(line string)
	Write() error
}

type EventParser interface {
	Parse(line string) (models.EventParsedDto, error)
}

type CompetitorService interface{}

var eventComments = map[uint16]string{
	1:  "The competitor(%d) registered",
	2:  "The start time for the competitor(%d) was set by a draw %s",
	3:  "The competitor(%d) is on the start line",
	4:  "The competitor(%d) has started",
	5:  "The competitor(%d) is on the firing range(%d)",
	6:  "The target(%d) has been hit by competitor(%d)",
	7:  "The competitor(%d) left the firing range",
	8:  "The competitor(%d) entered the penalty laps",
	9:  "The competitor(%d) left the penalty laps",
	10: "The competitor(%d) ended the main lap",
	11: "The competitor(%d) can`t continue: %s",
}

type EventConveyor struct {
	Parser EventParser
}

func NewEventConveyor(p EventParser) *EventConveyor {
	return &EventConveyor{Parser: p}
}

func (c *EventConveyor) ProcessEvents(events []string) error {
	// Пока что просто вывод с подстановкой значений
	for _, event := range events {
		eventParsed, err := c.Parser.Parse(event)
		if err != nil {
			return err
		}
		switch eventParsed.EventID {
		case 2:
			fmt.Printf(eventComments[eventParsed.EventID]+"\n",
				eventParsed.CompetitorID, eventParsed.ExtraParams["startTime"])
			break
		case 5:
			fmt.Printf(eventComments[eventParsed.EventID]+"\n",
				eventParsed.CompetitorID, eventParsed.ExtraParams["firingRange"])
			break
		case 6:
			fmt.Printf(eventComments[eventParsed.EventID]+"\n",
				eventParsed.ExtraParams["target"], eventParsed.CompetitorID)
			break
		case 11:
			fmt.Printf(eventComments[eventParsed.EventID]+"\n",
				eventParsed.CompetitorID, eventParsed.ExtraParams["comment"])
			break
		default:
			fmt.Printf(eventComments[eventParsed.EventID]+"\n", eventParsed.CompetitorID)
			break
		}
	}
	return nil
}
