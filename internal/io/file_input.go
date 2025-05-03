package io

import (
	"bufio"
	"os"

	"github.com/Dmitriy-M1319/biatlon-prototype/internal/conveyor"
)

type FileInputReaderImpl struct {
	Filename string
}

func NewFileInputReaderImpl(filename string) conveyor.EventsInputReader {
	return &FileInputReaderImpl{Filename: filename}
}

func (r *FileInputReaderImpl) ReadData() ([]string, error) {
	file, err := os.Open(r.Filename)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	result := make([]string, 0)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				return []string{}, err
			}
		}
		result = append(result, line[:len(line)-1])
	}
	return result, nil
}
