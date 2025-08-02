package main

import (
	"fmt"
	"net/http"

	"github.com/EuricoCruz/cleanarch_challenge/configs"
	"github.com/EuricoCruz/cleanarch_challenge/internal/infra/web"
)

func main() {
	configs, err := configs.LoadConfig("./"); if err != nil {
		panic(err)
	}

	webserver := webserver.NewWebServer(configs.WebServerPort)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	webserver.AddHandler("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Requisição recebida na raiz /")
		w.Write([]byte("Olá, mundo!"))
	})
	webserver.Start()
	fmt.Println("servidor rodando na porta: " + configs.WebServerPort)

}