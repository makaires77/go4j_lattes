package domain

// SearchResult representa os dados de um resultado de busca de currículo.
type SearchResult struct {
	Name string // Nome do pesquisador
	// Adicione aqui outros campos que você precisa extrair da busca, como e-mail, instituição, etc.
}

// SearchResults representa uma lista de resultados de busca de currículos.
type SearchResults []*SearchResult

// NewSearchResult cria uma nova instância de SearchResult.
func NewSearchResult(name string) *SearchResult {
	return &SearchResult{
		Name: name,
	}
}
