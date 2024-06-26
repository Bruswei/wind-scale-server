package csvdata

import (
	"encoding/csv"
	"fmt"
	"os"
	"wind-scale-server/internal/weatherservice"
)

type CSVStore struct {
	FilePath string
}

func NewCSVStore(filePath string) *CSVStore {
	return &CSVStore{
		FilePath: filePath,
	}
}

func (c *CSVStore) Create(record weatherservice.WindSpeedRecord) error {
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

func (c *CSVStore) Read() ([]weatherservice.WindSpeedRecord, error) {
	panic("Read method not implemented")
}

func (c *CSVStore) Update(record weatherservice.WindSpeedRecord) error {
	panic("Update method not implemented")
}

func (c *CSVStore) Delete(record weatherservice.WindSpeedRecord) error {
	panic("Delete method not implemented")
}
