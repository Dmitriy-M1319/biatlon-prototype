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
	t.Parallel()
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
	t.Parallel()
	p := parser.NewEventParserImpl()

	preparedString1 := "[ad:05:59.867] 1 1"
	preparedString2 := "[09:05:59.867] dsc 1"
	preparedString3 := "[ad:05:59.867] 1 aps"

	rslt, err := p.Parse(preparedString1)
	if err == nil {
		t.Error("Parse() doesn't find error", rslt)
	}

	rslt, err = p.Parse(preparedString2)
	if err == nil {
		t.Error("Parse() doesn't find error", rslt)
	}

	rslt, err = p.Parse(preparedString3)
	if err == nil {
		t.Error("Parse() doesn't find error", rslt)
	}
}

func TestParsingWithStartTime(t *testing.T) {
	t.Parallel()
	p := parser.NewEventParserImpl()
	preparedString := "[09:05:59.867] 2 1 09:30:00.000"
	tStmp, _ := time.Parse("15:04:05.000", "09:05:59.867")
	correctOutput := models.EventParsedDto{
		Timestamp:    tStmp,
		EventID:      2,
		CompetitorID: 1,
		ExtraParams:  map[string]string{"startTime": "09:30:00.000"},
	}

	funcOutput, err := p.Parse(preparedString)
	if err != nil {
		t.Error(err)
	}

	if !compareSimpleEventDtos(&funcOutput, &correctOutput) {
		t.Error("Event Dtos are not sample", funcOutput, correctOutput)
	}

	if funcOutput.ExtraParams["startTime"] != correctOutput.ExtraParams["startTime"] {
		t.Error("ExtraParams are not sample",
			funcOutput.ExtraParams["startTime"],
			"expected", correctOutput.ExtraParams["startTime"])
	}
}

func TestParsingWithFiringRange(t *testing.T) {
	t.Parallel()
	p := parser.NewEventParserImpl()
	preparedString := "[09:05:59.867] 5 1 1"
	tStmp, _ := time.Parse("15:04:05.000", "09:05:59.867")
	correctOutput := models.EventParsedDto{
		Timestamp:    tStmp,
		EventID:      5,
		CompetitorID: 1,
		ExtraParams:  map[string]string{"firingRange": "1"},
	}

	funcOutput, err := p.Parse(preparedString)
	if err != nil {
		t.Error(err)
	}

	if !compareSimpleEventDtos(&funcOutput, &correctOutput) {
		t.Error("Event Dtos are not sample", funcOutput, correctOutput)
	}

	if funcOutput.ExtraParams["firingRange"] != correctOutput.ExtraParams["firingRange"] {
		t.Error("ExtraParams are not sample",
			funcOutput.ExtraParams["firingRange"],
			"expected", correctOutput.ExtraParams["firingRange"])
	}
}

func TestParsingWithTarget(t *testing.T) {
	t.Parallel()
	p := parser.NewEventParserImpl()
	preparedString := "[09:05:59.867] 6 1 1"
	tStmp, _ := time.Parse("15:04:05.000", "09:05:59.867")
	correctOutput := models.EventParsedDto{
		Timestamp:    tStmp,
		EventID:      6,
		CompetitorID: 1,
		ExtraParams:  map[string]string{"target": "1"},
	}

	funcOutput, err := p.Parse(preparedString)
	if err != nil {
		t.Error(err)
	}

	if !compareSimpleEventDtos(&funcOutput, &correctOutput) {
		t.Error("Event Dtos are not sample", funcOutput, correctOutput)
	}

	if funcOutput.ExtraParams["target"] != correctOutput.ExtraParams["target"] {
		t.Error("ExtraParams are not sample",
			funcOutput.ExtraParams["target"],
			"expected", correctOutput.ExtraParams["target"])
	}
}

func TestParsingWithComment(t *testing.T) {
	t.Parallel()
	p := parser.NewEventParserImpl()
	preparedString := "[09:05:59.867] 11 1 Lost in the forest"
	tStmp, _ := time.Parse("15:04:05.000", "09:05:59.867")
	correctOutput := models.EventParsedDto{
		Timestamp:    tStmp,
		EventID:      11,
		CompetitorID: 1,
		ExtraParams:  map[string]string{"comment": "Lost in the forest"},
	}

	funcOutput, err := p.Parse(preparedString)
	if err != nil {
		t.Error(err)
	}

	if !compareSimpleEventDtos(&funcOutput, &correctOutput) {
		t.Error("Event Dtos are not sample", funcOutput, correctOutput)
	}

	if funcOutput.ExtraParams["comment"] != correctOutput.ExtraParams["comment"] {
		t.Error("ExtraParams are not sample",
			funcOutput.ExtraParams["comment"],
			"expected", correctOutput.ExtraParams["comment"])
	}
}
