package interfaces

import "github.com/your-user/go4j_lattes/pipeline/p1_find/domain"

// P1Interface define os métodos que a camada de aplicação irá chamar para interagir com a camada de interfaces.
type P1Interface interface {
	// ShowMessage exibe uma mensagem para o usuário.
	ShowMessage(message string)

	// ShowSearchResults exibe os resultados da busca de currículos para o usuário.
	ShowSearchResults(results domain.SearchResults)

	// GetSelectedName permite ao usuário escolher um currículo da lista de resultados.
	GetSelectedName(results domain.SearchResults) string
}
