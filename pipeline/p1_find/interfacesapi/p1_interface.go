package interfacesapi

import "github.com/makaires77/go4j_lattes/pipeline/p1_find/domain"

// CSVHandler define os métodos para lidar com o arquivo CSV.
type CSVHandler interface {
	// Métodos do CSVHandler aqui...
}

// QueueHandler define os métodos para lidar com a fila de pesquisadores.
type QueueHandler interface {
	// Métodos do QueueHandler aqui...
}

// SearchHandler define os métodos para lidar com a busca de pesquisadores.
type SearchHandler interface {
	// Métodos do SearchHandler aqui...
}

// P1Interface define os métodos que a camada de aplicação irá chamar para interagir com a camada de interfaces.
type P1Interface interface {
	// ShowMessage exibe uma mensagem para o usuário.
	ShowMessage(message string)

	// ShowSearchResults exibe os resultados da busca de currículos para o usuário.
	ShowSearchResults(results domain.SearchResults)

	// GetSelectedName permite ao usuário escolher um currículo da lista de resultados.
	GetSelectedName(results domain.SearchResults) string

	// FindResearchers realiza a busca de pesquisadores utilizando os handlers e interfaces fornecidos.
	FindResearchers(csvHandler CSVHandler, queueHandler QueueHandler, searchHandler SearchHandler) (domain.SearchResults, error)
}
