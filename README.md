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