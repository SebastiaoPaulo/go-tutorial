# Go Payment Integration API

Este projeto é uma API simples criada em Go como parte do meu aprendizado da linguagem. O objetivo é integrar a API da Stripe para gerenciar pagamentos e explorar conceitos como manipulação de rotas, decodificação de JSON e integração com serviços externos. Além disso, utilizei Docker e Docker Compose para facilitar a execução e distribuição da aplicação.

## Funcionalidades

- **Hello World Endpoint:** Retorna uma mensagem de texto simples.
- **Pagamento com Stripe:** Processa pagamentos com Stripe usando um `PaymentIntent`.

## Tecnologias Usadas

- [Go (Golang)](https://golang.org/)
- [Stripe Go SDK](https://github.com/stripe/stripe-go)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Requisitos

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Como Executar

1. Clone este repositório:

```bash
git clone https://github.com/usuario/seu-projeto.git
```

2. Entre no diretório do projeto:

```bash
cd seu-projeto
```

3. Execute o comando para rodar a aplicação usando Docker Compose:

```bash
docker-compose up
```

4. A aplicação estará disponível em `http://localhost:3333`.

## Endpoints

### 1. `GET /hello-world`

Retorna uma mensagem simples.

**Exemplo de Resposta:**

```text
Hello World!!!
```

### 2. `POST /payment`

Processa um pagamento com Stripe. Requer um JSON no corpo da requisição com o `productId`, `firstName`, e `lastName`.

**Exemplo de Request:**

```json
{
    "productId": "product111",
    "firstName": "John",
    "lastName": "Doe"
}
```

**Exemplo de Resposta:**

```json
{
    "clientSecret": "pi_1Abc123dEF45ghi6JKl789mno"
}
```

## Estrutura do Projeto

```bash
├── Dockerfile
├── docker-compose.yml
├── server.go
├── README.md
```

- **`server.go`**: Arquivo principal da aplicação.
- **`Dockerfile`**: Arquivo para construir a imagem Docker da aplicação.
- **`docker-compose.yml`**: Arquivo de configuração para rodar a aplicação com Docker Compose.

## Variáveis de Ambiente

- `STRIPE_API_KEY`: Chave secreta da API da Stripe. Essa chave está atualmente configurada diretamente no código (apenas para testes).

> **Nota:** Não se esqueça de substituir a chave da Stripe antes de usar em produção.

## Melhorias Futuras

- Adicionar autenticação.
- Implementar testes.
- Gerenciamento de produtos com um banco de dados.