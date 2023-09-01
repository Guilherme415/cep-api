# Cep API

A Cep API é uma aplicação que permite a busca de detalhes do endereço com base no CEP fornecido. Foi criada com o objetivo de possibilitar a busca de informações de endereços a partir de vários serviços diferentes de forma concorrente. O primeiro serviço que retornar as informações interromperá o processamento dos outros serviços, usando um contexto.

Caso um serviço não seja capaz de encontrar o endereço com base no CEP, ele realizará tentativas decrementando um dígito da direita para a esquerda até que o endereço seja localizado (Por exemplo: dado o CEP 22333999, serão feitas tentativas com 22333990, 22333900 e assim por diante).

## Arquitetura

A escolha da arquitetura para este projeto se baseia em padrões arquiteturais padronizados da linguagem de programação Go (Golang). Muitos desenvolvedores em Go provavelmente estão familiarizados com essa arquitetura ou com arquiteturas semelhantes. Essa escolha arquitetural permite incorporar facilmente serviços múltiplos, como workers, APIs REST, APIs gRPC e GraphQL. Tornam-se bastante simples ter vários serviços dentro desta aplicação, uma vez que a lógica de negócios fica separada da camada de apresentação, facilitando a reutilização de código. A arquitetura também se destaca por manter uma separação clara de responsabilidades. Por exemplo, as regras de negócios são organizadas no pacote `use_cases`, os serviços externos no pacote `services`, a configuração da aplicação em `config`, e assim por diante.

## Visão Geral do Processamento de Requisições

Aqui está um resumo dos passos envolvidos quando uma requisição é feita para a API:

1. Um cliente inicia uma requisição.
2. O cliente aciona uma chamada HTTP através de uma URL.
3. O cliente utiliza o DNS para encontrar o IP associado à URL.
4. Uma conexão TCP é estabelecida entre o cliente e o servidor.
5. O cliente especifica o tipo de requisição (GET, POST, etc.).
6. O servidor recebe e processa a requisição.
7. O servidor responde via conexão TCP, incluindo um código de status.
8. O cliente recebe a resposta.
9. A conexão TCP é encerrada.

## Configuração de Ambiente

Renomeie o arquivo `env.example` para `.env` e configure as variáveis de ambiente necessárias conforme sua preferência.
O token usado na API de busca de endereço está protegido por variável de ambiente

## Executando a Aplicação

Para executar a aplicação, siga os seguintes passos:

1. `go mod tidy`
2. `go run main.go`

Se desejar usar o Swagger para a documentação da API, execute o seguinte comando:

- `swag init`

Acesse a documentação do Swagger em:

- [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

Sinta-se à vontade para explorar a API e suas funcionalidades utilizando o Swagger UI fornecido.
