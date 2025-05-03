package io

import (
	"fmt"

	"github.com/Dmitriy-M1319/biatlon-prototype/internal/conveyor"
	"github.com/Dmitriy-M1319/biatlon-prototype/internal/models"
)

// TODO: Потом сделать просто прием отдельных сущностей и переводить в строки их уже здесь
var eventComments = map[uint16]string{
	1:  "[%s] The competitor(%d) registered",
	2:  "[%s] The start time for the competitor(%d) was set by a draw %s",
	3:  "[%s] The competitor(%d) is on the start line",
	4:  "[%s] The competitor(%d) has started",
	5:  "[%s] The competitor(%d) is on the firing range(%s)",
	6:  "[%s] The target(%s) has been hit by competitor(%d)",
	7:  "[%s] The competitor(%d) left the firing range",
	8:  "[%s] The competitor(%d) entered the penalty laps",
	9:  "[%s] The competitor(%d) left the penalty laps",
	10: "[%s] The competitor(%d) ended the main lap",
	11: "[%s] The competitor(%d) can`t continue: %s",
}

type ConsoleWriterImpl struct {
	Logs    []string
	Results []string
}

func NewConsoleWriterImpl() conveyor.OutputWriter {
	return &ConsoleWriterImpl{Logs: make([]string, 0), Results: make([]string, 0)}
}

func (w *ConsoleWriterImpl) InsertLog(e models.EventParsedDto) {
	var log string
	switch e.EventID {
	case 2:
		log = fmt.Sprintf(eventComments[e.EventID],
			e.Timestamp.Format("15:04:05.000"),
			e.CompetitorID, e.ExtraParams["startTime"])
		break
	case 5:
		log = fmt.Sprintf(eventComments[e.EventID],
			e.Timestamp.Format("15:04:05.000"),
			e.CompetitorID, e.ExtraParams["firingRange"])
		break
	case 6:
		log = fmt.Sprintf(eventComments[e.EventID],
			e.Timestamp.Format("15:04:05.000"),
			e.ExtraParams["target"], e.CompetitorID)
		break
	case 11:
		log = fmt.Sprintf(eventComments[e.EventID],
			e.Timestamp.Format("15:04:05.000"),
			e.CompetitorID, e.ExtraParams["comment"])
		break
	default:
		log = fmt.Sprintf(eventComments[e.EventID],
			e.Timestamp.Format("15:04:05.000"), e.CompetitorID)
		break
	}
	w.Logs = append(w.Logs, log)
}

func (w *ConsoleWriterImpl) InsertResult(line string) {
	w.Results = append(w.Results, line)
}

func (w *ConsoleWriterImpl) Write() error {
	fmt.Println("Output Log:")
	for _, log := range w.Logs {
		fmt.Println(log)
	}

	fmt.Println("")
	fmt.Println("Results:")
	for _, result := range w.Results {
		fmt.Println(result)
	}

	return nil
}
