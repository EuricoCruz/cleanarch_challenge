# Clean Arch Challenge

Este projeto faz parte do desafio da pós-graduação em Go do [Full Cycle](https://www.fullcycle.com.br/).

A aplicação consiste na criação e listagem de pedidos (orders) por meio de uma API REST, gRPC, e graphql. 

## 1. Clonar o repositório
git clone https://github.com/EuricoCruz/cleanarch_challenge

## 2. Levantar o container do banco de dados e do rabbitmq
docker-compose up -d

## 3. Criar tabela e inserir dados no banco de dados
O arquivo makefile já possui os comandos para criar a tabela, inserir dados e derrubar a tabela, caso necessário.

- Criar tabela
```bash
migrate create -ext=sql -dir=sql/migrations -seq init
```

- inserir dados
```bash
migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up
```

- remover tabeka e dados
```bash
migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose down
```


## 4. Baixar dependências
```bash
go mod tidy
````

## 5. Rodar a aplicação
```bash 
cd cmd/ordersystem 
go run main.go wire_gen.go 
```

# 5. Usando a aplicação

- 5.1 Criar pedido via API REST
```bash
curl --location --request POST 'localhost:8080/orders' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id":"bcde",
    "price": 100.5,
    "tax": 0.5
}'
```

- 5.2 Listar pedidos via API REST
```bash
curl --location --request GET 'localhost:8080/orders'
``` 

ou 

```bash
curl --location --request GET 'localhost:8080/orders?offsset=0&limit=20'
```

- 5.3 Criar pedido via gRPC
```bash
grpcurl -plaintext -d '{"id":"bcdx-e4fa","price":100.5,"tax":0.5}' localhost:50051 pb.OrderService/CreateOrder
``` 

- 5.4 Listar pedidos via gRPC
```bash
grpcurl -plaintext -d '{"offset":0,"limit":20}' localhost:50051 pb.OrderService/ListOrders
``` 

- 5.5 Criar pedido via GraphQL
```bash
curl --location 'localhost:8080/query' \
--header 'Content-Type: application/json' \
--data '{"query":"mutation {\n  createOrder(input: {id: \"bcde-1234\", Price: 103.5, Tax: 0.5}) {\n    id\n    Price\n    Tax\n    FinalPrice\n  }\n}\n"}'
``` 

- 5.6 Listar pedidos via GraphQL
```bash
curl --location 'localhost:8080/query' \
--header 'Content-Type: application/json' \
--data '{"query":"query {\n  listOrders(input: {offset: \"0\", limit: \"20\"}) {\n    id\n    Price\n    Tax\n    FinalPrice\n  }\n}\n"}'
``` 