# Mercado Bitcoin - Case Técnico: Variações de Médias Móveis Simples

Gerar um carregamento inicial para coletar a MMS dos últimos 365 dias e ter um endpoint GET para fornecer a MMS entre os dias 20, 50 e 200 de um range de dias

## Build

O processo de conteinerização do projeto é através da ferramenta <i>Docker Compose </i> onde um container da aplicação e um container do banco de dados Postgres são gerados ordenadamente. Isso facilita no processo de deploy, testes e resiliência da aplicação.

> Para buildar e executar é necessário apenas utilizar o comando do Makefile: <b>make start</b>

Esse comando inicializa a aplicação e o banco de dados.

## Tests

Nos testes são utilizados o framework [ginkgo](https://github.com/onsi/ginkgo), devido a falta de tempo, desenvolvi apenas testes unitários do serviço de coletar as MMS's mas a configuração para testes integrados já está preparada e pronta para ser utilizada.

> Para executar os testes é necessário apenas utilizar o comando do Makefile: <b>make test</b>

## API

A API utiliza apenas um endpoint para coletar as MMS's que são filtradas pelos seguintes parâmetros:

- FROM: (Timestamp) data de início do range de datas.
- TO: (Timestamp) data de fim do range de datas
- RANGE: (INT) range de dias para o cálculo da MMS. Valores fixos: (20, 50, 200) 
- Path: Utilizado 2 tipos de pares: BRLBTC e BRLETH

Exemplo de payload:

```
curl --location --request GET 'http://localhost:8080/BRLBTC/mms' \
--header 'Content-Type: application/json' \
--data '{
    "from": "2025-03-15T18:25:43.511Z",
    "to": "2025-03-18T18:25:43.511Z",
    "range": 20
}'

``` 

Exemplo de resposta do payload:

```
[
    {
        "timestamp": 1741996800,
        "mms": 529190.9986669752
    },
    {
        "timestamp": 1742083200,
        "mms": 525409.0475227868
    },
    {
        "timestamp": 1742169600,
        "mms": 522814.1941368308
    },
    {
        "timestamp": 1742256000,
        "mms": 520758.41681296536
    }
]
```

## Script

A solução que adotei foi criar um serviço que ao iniciar a aplicação irá realizadas duas chamadas provenientes dos pares solicitados (BRLBTC, BRLETH) na API do mercado bitcoin, coletar os dados de 1 ano e 200 dias atrás e a partir desses dados fazer o cálculo baseado no range de dias e a conversão após isso realizar a persistência em uma tabela do banco de dados.

## Job

Uma estratégia para incrementar diariamente na tabela seria criar um worker serverless que irá executar em determinada hora fazendo o cálculo referentes aos dias padrões de MMS e persistindo na tabela.