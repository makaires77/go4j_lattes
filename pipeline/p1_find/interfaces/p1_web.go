// pipeline/p1_find/interfaces/p1_web.go
package interfaces

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/makaires77/go4j_lattes/pipeline/p1_find/domain"
)

// P1Web é a implementação da interface P1Interface para interação com o usuário via CLI.
type P1Web struct {
	scanner *bufio.Scanner
}

// NewP1Web cria uma nova instância de P1Web.
func NewP1Web() *P1Web {
	return &P1Web{
		scanner: bufio.NewScanner(os.Stdin),
	}
}

// ShowMessage exibe uma mensagem para o usuário.
func (web *P1Web) ShowMessage(message string) {
	fmt.Println(message)
}

// ShowSearchResults exibe os resultados da busca de currículos para o usuário.
func (web *P1Web) ShowSearchResults(results domain.SearchResults) {
	fmt.Println("Resultados da busca de currículos:")
	for i, result := range results {
		fmt.Printf("%d. %s\n", i+1, result.Name)
	}
}

// GetSelectedName permite ao usuário escolher um currículo da lista de resultados.
func (web *P1Web) GetSelectedName(results domain.SearchResults) string {
	for {
		fmt.Print("Digite o número do currículo desejado ou 's' para sair: ")
		web.scanner.Scan()
		input := strings.TrimSpace(web.scanner.Text())

		if input == "s" {
			os.Exit(0)
		}

		num, err := strconv.Atoi(input)
		if err != nil || num < 1 || num > len(results) {
			fmt.Println("Opção inválida. Digite novamente.")
			continue
		}

		// O usuário escolheu um currículo válido
		selectedName := results[num-1].Name
		return selectedName
	}
}
