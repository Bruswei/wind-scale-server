package csvdata

import (
	"encoding/csv"
	"fmt"
	"os"
	"wind-scale-server/internal/windspeed"
)

type CSVStore struct {
	FilePath string
}

func NewCSVStore(filePath string) *CSVStore {
	return &CSVStore{
		FilePath: filePath,
	}
}

func (c *CSVStore) StoreData(record windspeed.WindSpeedRecord) error {
	file, err := os.OpenFile(c.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	line := []string{record.Location, record.Time, fmt.Sprintf("%f", record.WindSpeed)}
	if err := writer.Write(line); err != nil {
		return err
	}

	return nil
}
