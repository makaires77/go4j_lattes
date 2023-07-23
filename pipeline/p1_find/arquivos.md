 # Funcionalidade de cada arquivo

### Pasta 'pipeline/p1_find/application.'
**csv_handler.go**: função responsável por manipular o arquivo CSV e obter a lista de nomes para realizar a busca. Adicione também funções auxiliares para tratar erros e interagir com o usuário.

**queue_handler.go**: função responsável por manipular a fila de nomes para processamento em paralelo. Também adicionaremos funções auxiliares para adicionar nomes à fila, obter nomes da fila e verificar se a fila está vazia.

**search_handler.go**: função responsável por realizar as buscas dos currículos na Plataforma Lattes do CNPq. Adicionaremos funções auxiliares para executar a busca, verificar se houve erro na requisição e realizar novas tentativas em caso de falhas.


### Pasta 'pipeline/p1_find/cmd'
**p1_cmd.go**: função que será responsável por gerenciar a execução do processo de busca dos currículos na Plataforma Lattes do CNPq. Adicionaremos funções auxiliares para interagir com o usuário, processar a lista de nomes do arquivo CSV e executar as buscas em paralelo. 

### Pasta 'pipeline/p1_find/domain'
**p1_domain.go**: Contém as estruturas de dados que serão usadas para representar as entidades relacionadas à busca de currículos. Também adicionaremos funções auxiliares, se necessário, para realizar validações ou manipulações dos dados. 

### Pasta 'pipeline/p1_find/infra'
**p1_infra.go**: função que será responsável por interagir com a infraestrutura para armazenar os dados dos currículos. Neste caso, usaremos a infraestrutura para armazenar os resultados das buscas em um banco de dados, mas você pode modificar de acordo com a infraestrutura que estiver usando. criamos a estrutura P1Infra, que é responsável por interagir com a infraestrutura para armazenar os dados dos currículos. A função *SaveSearchResult* é responsável por salvar o resultado da busca de um currículo na infraestrutura. Aqui, estamos apenas imprimindo uma mensagem no console para simular o salvamento dos dados, antes de implementar a lógica adequada para salvar os dados na infraestrutura que você estiver usando.

Observe que no momento estamos apenas simulando o salvamento dos dados para fins de demonstração. Você deve substituir a implementação da função SaveSearchResult para se adequar à infraestrutura real que você está utilizando. 

### Pasta 'pipeline/p1_find/interfaces'
**p1_infra.go**: interface que representa o contrato entre a camada de aplicação e a camada de interfaces. Esta interface irá definir os métodos que a camada de aplicação irá chamar para interagir com a camada de interfaces, sem depender diretamente dela. Criamos a interface P1Interface, que define três métodos que a camada de aplicação irá chamar para interagir com a camada de interfaces:

    ShowMessage: Exibe uma mensagem para o usuário.

    ShowSearchResults: Exibe os resultados da busca de currículos para o usuário.
    
    GetSelectedName: Permite ao usuário escolher um currículo da lista de resultados.

Essa interface define um contrato que será implementado pela camada de interfaces. Dessa forma, a camada de aplicação pode chamar esses métodos sem depender diretamente da implementação específica da interface. Isso permite que você substitua facilmente a camada de interfaces por outra implementação, se necessário, sem afetar a camada de aplicação.

**p1_web.go**: implementação da interface P1Interface para a interação com o usuário através de uma interface de linha de comando (CLI). Vamos utilizar a biblioteca fmt para exibir as mensagens e ler a entrada do usuário. P1Web, que é a implementação da interface P1Interface para interagir com o usuário via CLI (linha de comando). Utilizamos a biblioteca bufio para criar um scanner para ler a entrada do usuário a partir do os.Stdin.

    A função ShowMessage exibe uma mensagem para o usuário na saída padrão.

    A função ShowSearchResults exibe os resultados da busca de currículos para o usuário, apresentando a lista numerada dos nomes dos pesquisadores.

    A função GetSelectedName permite ao usuário escolher um currículo da lista de resultados digitando o número correspondente ou 's' para sair. A função lê a entrada do usuário, verifica se é um número válido e corresponde a um currículo da lista, e retorna o nome do currículo selecionado.

Observe que esta é apenas uma implementação básica para interação via CLI, e você pode personalizar a interação conforme necessário para se adequar ao seu projeto.


