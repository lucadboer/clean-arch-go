# Desafio de implementação - Clean Arch

Este desafio envolve a criação de um usecase para listagem de orders utilizando diferentes abordagens: REST, gRPC e GraphQL.

## Pré-requisitos

- Docker e Docker Compose instalados.
- Go (Golang) instalado para desenvolvimento do backend.

## Configuração do Ambiente

1. **Banco de Dados e RabbitMQ**: Utilize o `docker-compose.yaml` fornecido para subir o MySQL e o RabbitMQ. Execute o comando abaixo para iniciar os serviços:

```bash
docker-compose up -d
```
````

2. **Migrações**: Crie as migrações necessárias para a tabela `orders` no banco de dados MySQL. Você pode usar uma ferramenta como `Flyway`, `GORM` ou outra de sua preferência para gerenciar as migrações.

Exemplo de migração para criar a tabela `orders`:

```sql
CREATE TABLE orders (
    id VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    tax DECIMAL(10, 2) NOT NULL,
    final_price DECIMAL(10, 2) NOT NULL,
    PRIMARY KEY (id)
);
```

3. **Desenvolvimento**:

- **REST Endpoint**: Crie um endpoint GET `/order` para listar todas as orders.
- **Service ListOrders com GRPC**: Implemente um serviço gRPC para listar as orders.
- **Query ListOrders GraphQL**: Implemente uma query GraphQL para listar as orders.

4. **Arquivo `order.http`**: Inclua um arquivo `order.http` com requests de exemplo para criar e listar orders. Exemplo:

```http
POST http://localhost:8000/order HTTP/1.1
Host: localhost:8000
Content-Type: application/json

{
    "id":"order1",
    "price": 100.5,
    "tax": 0.5
}

###
GET http://localhost:8000/order HTTP/1.1
Host: localhost:8000
```

## Executando a Aplicação

- **Web Server**: Inicie o servidor web na porta 8000.
- **gRPC Server**: Inicie o servidor gRPC na porta 50051.
- **GraphQL Server**: Inicie o servidor GraphQL na porta 8080.

## Testando a Aplicação

Após iniciar todos os serviços, você pode testar cada um utilizando o arquivo `order.http` para o endpoint REST, ferramentas como `grpcurl` ou `evans` para o gRPC e um cliente GraphQL como `Insomnia` ou `Postman` para GraphQL, ou entrando no próprio http://localhost:8080 para realizar as querys e mutates.
