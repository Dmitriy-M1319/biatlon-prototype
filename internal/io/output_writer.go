package io

import (
	"fmt"

	"github.com/Dmitriy-M1319/biatlon-prototype/internal/config"
	"github.com/Dmitriy-M1319/biatlon-prototype/internal/conveyor"
	"github.com/Dmitriy-M1319/biatlon-prototype/internal/models"
)

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

func (w *ConsoleWriterImpl) InsertResult(r *models.CompetitorResult) {
	var result string
	if len(r.Status) > 0 {
		result = fmt.Sprintf("[%s] %d", r.Status, r.CompetitorID)
	} else {
		result = fmt.Sprintf("[%02d:%02d:%02d.%03d] %d", int(r.DistanceTime.Hours()),
			int(r.DistanceTime.Minutes())%60, int(r.DistanceTime.Seconds())%60,
			int(r.DistanceTime.Milliseconds())%1000, r.CompetitorID)
	}

	startMainLaps := "["
	for i, lap := range r.MainLaps {
		startMainLaps += fmt.Sprintf("{%02d:%02d:%02d.%03d, %.3f}", int(lap.Duration.Hours()),
			int(lap.Duration.Minutes())%60, int(lap.Duration.Seconds())%60,
			int(lap.Duration.Milliseconds())%1000, lap.Speed)
		if i != len(r.MainLaps)-1 {
			startMainLaps += ", "
		}
	}

	if len(r.MainLaps) < config.Config().Laps {
		for _ = range config.Config().Laps - len(r.MainLaps) {
			startMainLaps += ", {,}"
		}
	}
	startMainLaps += "]"
	result += " " + startMainLaps

	if r.PenaltyAvgSpeed > 0.0 {
		result += fmt.Sprintf(" {%02d:%02d:%02d.%03d, %.3f}", int(r.PenaltyTime.Hours()),
			int(r.PenaltyTime.Minutes())%60, int(r.PenaltyTime.Seconds())%60,
			int(r.PenaltyTime.Milliseconds())%1000, r.PenaltyAvgSpeed)
	} else {
		result += " {,}"
	}

	result += fmt.Sprintf(" %d/%d", r.TotalHits, r.TotalShots)

	w.Results = append(w.Results, result)
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
