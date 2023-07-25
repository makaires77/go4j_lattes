// No arquivo p1_find/application/csv_handler.go

package application

import (
	"encoding/csv"
	"os"
)

type CSVHandler struct {
	FilePath string
}

func NewCSVHandler(filePath string) *CSVHandler {
	return &CSVHandler{
		FilePath: filePath,
	}
}

func (c *CSVHandler) ReadCSV() ([]string, error) {
	file, err := os.Open(c.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var names []string
	for _, line := range lines {
		if len(line) > 0 {
			names = append(names, line[0])
		}
	}

	return names, nil
}
