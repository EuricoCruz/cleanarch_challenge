# Clean Arch Challenge

Este projeto faz parte do desafio da pós-graduação em Go do [Full Cycle](https://www.fullcycle.com.br/).

A aplicação consiste na criação e listagem de pedidos (orders) por meio de uma API REST, gRPC, e graphql. 

## 1. Clonar o repositório
git clone https://github.com/EuricoCruz/cleanarch_challenge

## 2. Levantar o container da aplicação e de suas dependências
docker-compose up -d


## 3. Usando a aplicação

Com a aplicação rodando os arquivos create_order.http e list_order.http na pasta api podem ser usados no [Insomnia](https://insomnia.rest/) ou [Postman](https://www.postman.com/) para testar a aplicação via REST, GRPC e GraphQL.

P.S: a aplicação demora um pouco para ficar totalmente disponível, pois precisa aguardar o banco de dados e o RabbitMQ ficarem prontos.

As aplicações estarão disponíveis nas seguintes portas:
- API REST: http://localhost:8000
- gRPC: localhost:50051
- GraphQL: http://localhost:8080/query

Além disso, seguem alguns exemplos de comandos curl e grpcurl para testar a aplicação via terminal.

- 3.1 Criar pedido via API REST
```bash
curl --location --request POST 'localhost:8080/orders' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id":"bcde",
    "price": 100.5,
    "tax": 0.5
}'
```

- 3.2 Listar pedidos via API REST
```bash
curl --location --request GET 'localhost:8080/orders'
``` 

ou 

```bash
curl --location --request GET 'localhost:8080/orders?offsset=0&limit=20'
```

- 3.3 Criar pedido via gRPC
```bash
grpcurl -plaintext -d '{"id":"bcdx-e4fa","price":100.5,"tax":0.5}' localhost:50051 pb.OrderService/CreateOrder
``` 

- 3.4 Listar pedidos via gRPC
```bash
grpcurl -plaintext -d '{"offset":0,"limit":20}' localhost:50051 pb.OrderService/ListOrders
``` 

- 3.5 Criar pedido via GraphQL
```bash
curl --location 'localhost:8080/query' \
--header 'Content-Type: application/json' \
--data '{"query":"mutation {\n  createOrder(input: {id: \"bcde-1234\", Price: 103.5, Tax: 0.5}) {\n    id\n    Price\n    Tax\n    FinalPrice\n  }\n}\n"}'
``` 

- 3.6 Listar pedidos via GraphQL
```bash
curl --location 'localhost:8080/query' \
--header 'Content-Type: application/json' \
--data '{"query":"query {\n  listOrders(input: {offset: \"0\", limit: \"20\"}) {\n    id\n    Price\n    Tax\n    FinalPrice\n  }\n}\n"}'
``` 