FROM golang:1.25.0-alpine3.21

WORKDIR /app

COPY . ./
RUN go mod download

CMD ["go", "run", "cmd/ordersystem/wire_gen.go", "cmd/ordersystem/main.go"]