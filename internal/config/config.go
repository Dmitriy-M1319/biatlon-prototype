package config

import (
	"encoding/json"
	"os"

	"github.com/Dmitriy-M1319/biatlon-prototype/internal/models"
)

// TODO: Тесты

var config *models.Config

func Config() *models.Config {
	return config
}

func ParseJsonConfig(data []byte) error {
	result := models.Config{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		return err
	}

	config = &result
	return nil
}

func ReadFileConfigJson(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return ParseJsonConfig(data)
}
