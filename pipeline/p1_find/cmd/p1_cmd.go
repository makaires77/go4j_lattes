// pipeline/p1_find/cmd/p1_cmd.go
package cmd

import (
	"fmt"
	"sync"

	"github.com/makaires77/go4j_lattes/pipeline/p1_find/application"
)

// P1Cmd é responsável por gerenciar a execução do processo de busca dos currículos.
type P1Cmd struct {
	CSVFilePath string // Caminho do arquivo CSV com a lista de nomes
}

// NewP1Cmd cria uma nova instância de P1Cmd.
func NewP1Cmd(csvFilePath string) *P1Cmd {
	return &P1Cmd{
		CSVFilePath: csvFilePath,
	}
}

// Execute inicia o processo de busca dos currículos.
func (cmd *P1Cmd) Execute() error {
	// Criar instâncias dos handlers
	csvHandler := application.NewCSVHandler(cmd.CSVFilePath)
	queueHandler := application.NewQueueHandler()
	searchHandler := application.NewSearchHandler()

	// Obter a lista de nomes do arquivo CSV
	names, err := csvHandler.GetNamesFromCSV()
	if err != nil {
		return fmt.Errorf("erro ao obter a lista de nomes do arquivo CSV: %w", err)
	}

	// Adicionar os nomes à fila para processamento em paralelo
	for _, name := range names {
		queueHandler.AddToQueue(name)
	}

	// Definir a quantidade de workers para processar os nomes em paralelo
	numWorkers := 10

	// Processar a fila em paralelo
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				name, ok := queueHandler.GetFromQueue()
				if !ok {
					return
				}

				err := searchHandler.PerformSearch(name)
				if err != nil {
					fmt.Printf("Erro na busca do nome \"%s\": %v\n", name, err)
				}
			}
		}()
	}

	wg.Wait()

	return nil
}
