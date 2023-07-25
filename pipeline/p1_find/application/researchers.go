// p1_find/application/researchers.go
package application

import (
	"fmt"
	"strings"

	"github.com/makaires77/go4j_lattes/pipeline/p1_find/domain"
)

func FindResearchers(csvHandler CSVHandler, queueHandler QueueHandler, searchHandler SearchHandler) ([]domain.Researcher, error) {
	// Ler os nomes do arquivo CSV usando o CSVHandler
	names, err := csvHandler.ReadCSV()
	if err != nil {
		return nil, fmt.Errorf("failed to read names from CSV: %w", err)
	}

	var researchers []domain.Researcher

	// Para cada nome na lista, realizar a busca usando o SearchHandler
	for _, name := range names {
		// Remover espaços em branco do nome
		name = strings.TrimSpace(name)

		if name != "" {
			// Realizar a busca usando o SearchHandler
			result, err := searchHandler.FindResearchers(name)
			if err != nil {
				return nil, fmt.Errorf("failed to find researchers: %w", err)
			}

			// Adicionar os resultados à lista de pesquisadores
			researchers = append(researchers, result...)
		}
	}

	// Salvar os resultados na fila usando o QueueHandler
	err = queueHandler.SaveResults(researchers)
	if err != nil {
		return nil, fmt.Errorf("failed to save results to queue: %w", err)
	}

	return researchers, nil
}
