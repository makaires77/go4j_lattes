// pipeline\p1_find\application\search_handler.go
package application

import (
	"fmt"
	"net/http"
	"time"
)

// SearchHandler é responsável por realizar as buscas dos currículos na Plataforma Lattes do CNPq.
type SearchHandler struct {
}

// NewSearchHandler cria uma nova instância de SearchHandler.
func NewSearchHandler() *SearchHandler {
	return &SearchHandler{}
}

// PerformSearch realiza a busca de um nome na Plataforma Lattes do CNPq.
func (handler *SearchHandler) PerformSearch(name string) error {
	url := "http://buscatextual.cnpq.br/buscatextual/busca.do"

	// Construir a requisição POST com o nome a ser buscado
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return fmt.Errorf("não foi possível construir a requisição: %w", err)
	}

	q := req.URL.Query()
	q.Add("buscarDemais", "true")
	q.Add("textoBusca", name)
	req.URL.RawQuery = q.Encode()

	// Executar a requisição
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("não foi possível realizar a busca: %w", err)
	}
	defer resp.Body.Close()

	// Verificar se houve erro na requisição
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("busca retornou status %d", resp.StatusCode)
	}

	// Processar a resposta da busca e extrair os dados dos currículos encontrados
	// Implemente aqui a lógica para extrair os dados dos currículos

	return nil
}
