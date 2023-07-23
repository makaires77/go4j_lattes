package tests

import (
	"testing"

	"github.com/makaires77/go4j_lattes/pipeline/p1_find/application"
	"github.com/makaires77/go4j_lattes/pipeline/p1_find/domain"
	"github.com/makaires77/go4j_lattes/pipeline/p1_find/interfaces"

	"github.com/stretchr/testify/assert"
)

func TestP1(t *testing.T) {
	// Simulando resultados da busca
	results := domain.SearchResults{
		{
			Name: "Raimir Holanda Filho",
			// Adicione outros campos que você queira simular
		},
		{
			Name: "Nabor das Chagas Mendonça",
			// Adicione outros campos que você queira simular
		},
		// Adicione mais resultados simulados conforme necessário
	}

	// Criando as instâncias dos componentes
	csvHandler := application.NewCSVHandler("")
	queueHandler := application.NewQueueHandler()
	searchHandler := application.NewSearchHandler()
	webInterface := interfaces.NewP1Web()

	// Testando a função FindResearchers
	t.Run("Teste FindResearchers", func(t *testing.T) {
		// Chame a função FindResearchers com os parâmetros desejados e verifique se os resultados são os esperados
		researchers, err := application.FindResearchers(csvHandler, queueHandler, searchHandler, webInterface)
		assert.NoError(t, err)
		assert.Equal(t, results, researchers)
	})
}
