package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Dmitriy-M1319/biatlon-prototype/internal/conveyor"
	"github.com/Dmitriy-M1319/biatlon-prototype/internal/io"
	"github.com/Dmitriy-M1319/biatlon-prototype/internal/parser"
)

func main() {
	eventFile := flag.String("eventsFile", "", "File with incoming events")
	flag.Parse()

	if *eventFile == "" {
		fmt.Println("Please provide a filename using the -eventsFile flag.")
		return
	}

	r := io.NewFileInputReaderImpl(*eventFile)
	w := io.NewConsoleWriterImpl()
	pSer := parser.NewEventParserImpl()
	conv := conveyor.NewEventConveyor(pSer, w, r)

	err := conv.StartProcessEvents()
	if err != nil {
		log.Fatal(err)
	}
}
