package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/EuricoCruz/cleanarch_challenge/configs"
	"github.com/EuricoCruz/cleanarch_challenge/internal/event/handler"
	"github.com/EuricoCruz/cleanarch_challenge/internal/infra/web/webserver"
	"github.com/EuricoCruz/cleanarch_challenge/pkg/events"
	"github.com/streadway/amqp"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)
func main() {
	configs, err := configs.LoadConfig("."); if err != nil {
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

	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	webserver.Post("/order", webOrderHandler.Create)
	webserver.Get("/order", webOrderHandler.List)
	webserver.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Requisição recebida na raiz /")
		w.Write([]byte("Olá, mundo!"))
	})
	webserver.Start()
	fmt.Println("servidor rodando na porta: " + configs.WebServerPort)

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