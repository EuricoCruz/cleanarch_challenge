package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/EuricoCruz/cleanarch_challenge/configs"
	"github.com/EuricoCruz/cleanarch_challenge/internal/event/handler"
	"github.com/EuricoCruz/cleanarch_challenge/internal/infra/graph"
	"github.com/EuricoCruz/cleanarch_challenge/internal/infra/grpc/pb"
	"github.com/EuricoCruz/cleanarch_challenge/internal/infra/grpc/service"
	"github.com/EuricoCruz/cleanarch_challenge/internal/infra/web/webserver"
	"github.com/EuricoCruz/cleanarch_challenge/pkg/events"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rabbitMQChannel := getRabbitMQChannel()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)
	listOrderUseCase := NewListOrderUseCase(db)

	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	webserver.Post("/order", webOrderHandler.Create)
	webserver.Get("/order", webOrderHandler.List)
	webserver.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Requisição recebida na raiz /")
		w.Write([]byte("Olá, mundo!"))
	})
	go webserver.Start()
	fmt.Println("servidor rodando na porta: " + configs.WebServerPort)

	grpcServer := grpc.NewServer()
	orderService := service.NewOrderService(*createOrderUseCase, *listOrderUseCase)
	pb.RegisterOrderServiceServer(grpcServer, orderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}

	go func() {
		fmt.Println("Servidor gRPC rodando na porta", configs.GRPCServerPort)
		if err := grpcServer.Serve(lis); err != nil {
			panic(err)
		}
	}()

	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
		ListOrderUseCase:   *listOrderUseCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
	http.ListenAndServe(":"+configs.GraphQLServerPort, nil)

}

func getRabbitMQChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}
