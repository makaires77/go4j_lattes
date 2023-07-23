package application

import (
	"encoding/csv"
	"fmt"
	"os"
)

// CSVHandler é responsável por manipular o arquivo CSV e obter a lista de nomes.
type CSVHandler struct {
	Filepath string // Caminho do arquivo CSV
}

// NewCSVHandler cria uma nova instância de CSVHandler.
func NewCSVHandler(filepath string) *CSVHandler {
	return &CSVHandler{
		Filepath: filepath,
	}
}

// GetNamesFromCSV lê o arquivo CSV e retorna a lista de nomes.
func (handler *CSVHandler) GetNamesFromCSV() ([]string, error) {
	// Abrir o arquivo CSV
	file, err := os.Open(handler.Filepath)
	if err != nil {
		return nil, fmt.Errorf("não foi possível abrir o arquivo CSV: %w", err)
	}
	defer file.Close()

	// Criar um leitor CSV
	reader := csv.NewReader(file)

	// Ler os registros do arquivo CSV
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("não foi possível ler os registros do arquivo CSV: %w", err)
	}

	// Extrair a lista de nomes da primeira coluna do CSV
	var names []string
	for _, record := range records {
		if len(record) > 0 {
			names = append(names, record[0])
		}
	}

	return names, nil
}
