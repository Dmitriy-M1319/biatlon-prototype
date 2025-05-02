package parser_test

import (
	"testing"
	"time"

	"github.com/Dmitriy-M1319/biatlon-prototype/internal/models"
	"github.com/Dmitriy-M1319/biatlon-prototype/internal/parser"
)

func compareSimpleEventDtos(d1 *models.EventParsedDto, d2 *models.EventParsedDto) bool {
	return d1.Timestamp == d2.Timestamp &&
		d1.EventID == d2.EventID &&
		d1.CompetitorID == d1.CompetitorID
}

func TestSimpleEventCorrectParsing(t *testing.T) {
	p := parser.NewEventParserImpl()

	preparedString := "[09:05:59.867] 1 1"
	tStmp, _ := time.Parse("15:04:05.000", "09:05:59.867")
	correctOutput := models.EventParsedDto{
		Timestamp:    tStmp,
		EventID:      1,
		CompetitorID: 1,
	}

	funcOutput, err := p.Parse(preparedString)
	if err != nil {
		t.Error(err)
	}

	if !compareSimpleEventDtos(&funcOutput, &correctOutput) {
		t.Error("Event Dtos are not sample", funcOutput, correctOutput)
	}
}

func TestSimpleEventInvalidParsing(t *testing.T) {
	t.Error("Not implemented")
}

// TODO: Оставшиеся тесты сделать с параллельным запуском (один - успешный парсинг, второй - невалидный)

func TestParsingWithStartTime(t *testing.T) {
	t.Error("Not implemented")
}

func TestParsingWithFiringRange(t *testing.T) {
	t.Error("Not implemented")
}

func TestParsingWithTarget(t *testing.T) {
	t.Error("Not implemented")
}

func TestParsingWithComment(t *testing.T) {
	t.Error("Not implemented")
}
