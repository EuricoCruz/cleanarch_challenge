package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/EuricoCruz/cleanarch_challenge/configs"
	"github.com/EuricoCruz/cleanarch_challenge/internal/infra/web"
	
	// mysql
	_ "github.com/go-sql-driver/mysql"
)
func main() {
	configs, err := configs.LoadConfig("./"); if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	webserver := webserver.NewWebServer(configs.WebServerPort)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	webserver.AddHandler("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Requisição recebida na raiz /")
		w.Write([]byte("Olá, mundo!"))
	})
	webserver.Start()
	fmt.Println("servidor rodando na porta: " + configs.WebServerPort)

}