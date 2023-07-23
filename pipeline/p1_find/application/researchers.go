package application

import (
	"github.com/makaires77/go4j_lattes/pipeline/p1_find/domain"
	"github.com/makaires77/go4j_lattes/pipeline/p1_find/interfaces"
)

// FindResearchers é responsável por encontrar os pesquisadores utilizando os handlers e interfaces fornecidos.
func (handler *SearchHandler) FindResearchers(
	csvHandler *CSVHandler,
	queueHandler *QueueHandler,
	searchHandler *SearchHandler,
	webInterface interfaces.P1Web,
) (domain.SearchResults, error) {
	// Implemente a lógica da função FindResearchers aqui, utilizando os handlers e interfaces fornecidos.
	// ...

	// Exemplo de utilização da nova interface P1Web
	webInterface.ShowMessage("Pesquisa de currículos na Plataforma Lattes do CNPq")
	webInterface.ShowMessage("Digite o nome do pesquisador a ser buscado:")
	// ...

	return nil, nil
}
