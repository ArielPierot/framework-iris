# Iris Framework Project

O objetivo desse projeto é utilizar o Iris Framework na construção de um Rest API com Basic Auth.

# Começando

Siga as instruções a seguir para executar o projeto em sua máquina para testes.


# Pré-requisitos

Foi utilizado nesse projeto o Go na versão 1.16 para construção desse projeto.

# Deployment

É possível rodar de qualquer máquina que tenha o docker instalado com o comando a seguir:

```bash
    docker compose up -d --build
```

# API

## Autorização

A API usa basic authentication. É possível adicionar outros usuários no arquivo `users.yml`. O hash da senha é gerada usando Bcrypt. Com ela é possível acessar todos os endpoints da aplicação.

### Token

`dGVzdDoxMjM0NTY=`

## Recursos

### produto

- **POST** /produto

Criar novo produto

```bash
    curl --request POST \
    --url http://localhost:8080/produto \
    --header 'Authorization: Basic dGVzdDoxMjM0NTY=' \
    --header 'Content-Type: application/json' \
    --data '{
            "codigo": "1234",
            "nome": "Produto B",
            "preco_de": 23.99,
            "preco_por": 19.99,
            "estoque": {
                "estoque_total": 300,
                "estoque_corte": 25		
            }	
        }'
```

- **GET** /produto/**\<codigo\>**

Recupera informações do produto através do código

```bash
    curl --request GET \
    --url http://localhost:8080/produto/1234 \
    --header 'Authorization: Basic dGVzdDoxMjM0NTY='
```

- **GET** /produto/?page=**\<page number\>**&page_size=**\<page size\>**

Lista de produtos com paginação

```bash
    curl --request GET \
    --url 'http://localhost:8080/produto?page=1&page_size=2' \
    --header 'Authorization: Basic dGVzdDoxMjM0NTY='
```

- **PUT** /produto/**\<codigo\>**

Atualiza as informações do produto

```bash
    curl --request PUT \
    --url http://localhost:8080/produto/1234 \
    --header 'Authorization: Basic dGVzdDoxMjM0NTY=' \
    --header 'Content-Type: application/json' \
    --data '{
        "nome": "Produto C",
        "preco_de": 21.99,
        "preco_por": 17.99,
        "estoque": {
            "estoque_total": 300,
            "estoque_corte": 30
        }	
    }'
```

- **DELETE** /produto/**\<codigo\>**

Remove o produto da base de dados

```bash
    curl --request DELETE \
    --url http://localhost:8080/produto/1234 \
    --header 'Authorization: Basic dGVzdDoxMjM0NTY='
```