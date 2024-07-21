# Desafio de Multithreading e APIs

Neste desafio, o objetivo é utilizar conceitos de multithreading e realizar requisições simultâneas a duas APIs de consulta de CEP. A ideia é identificar qual API responde mais rapidamente e descartar a resposta da API mais lenta.

## APIs Utilizadas

As duas APIs que serão utilizadas são:

1. **BrasilAPI**: Uma API brasileira de consulta de CEP.
   - Endpoint: `https://brasilapi.com.br/api/cep/v1/{cep}`

2. **ViaCEP**: Uma API brasileira de consulta de CEP.
   - Endpoint: `http://viacep.com.br/ws/{cep}/json/`

## Requisitos do Desafio

- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.
- O resultado da request deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou.
- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.

## Como Executar

Certifique-se de ter Go instalado na sua máquina. Você pode instalá-lo através do site oficial: https://golang.org/

Para executar o código deste desafio em Go:

1. Clone este repositório para sua máquina local.
2. Navegue até o diretório onde o código foi clonado.
3. Execute o seguinte comando para compilar e executar o código:

```bash
go run server/main.go