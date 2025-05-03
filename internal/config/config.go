package config

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Dmitriy-M1319/biatlon-prototype/internal/models"
)

// TODO: Тесты

var config *models.Config

func Config() *models.Config {
	return config
}

func parseDuration(s string) (time.Duration, error) {
	parts := strings.Split(s, ":")
	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}

	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, err
	}

	seconds, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, err
	}

	duration := time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute + time.Duration(seconds)*time.Second
	return duration, nil
}

func ParseJsonConfig(data []byte) error {
	result := models.Config{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		return err
	}

	result.StartDeltaTime, err = parseDuration(result.StartDeltaStr)
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
