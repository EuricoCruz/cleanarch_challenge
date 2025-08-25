package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	return &WebServer{
		Router:        router,
		WebServerPort: serverPort,
	}
}

func (s *WebServer) Get(path string, handler http.HandlerFunc) {
	s.Router.Get(path, handler)
}

func (s *WebServer) Post(path string, handler http.HandlerFunc) {
	s.Router.Post(path, handler)
}

func (s *WebServer) Put(path string, handler http.HandlerFunc) {
	s.Router.Put(path, handler)
}

func (s *WebServer) Delete(path string, handler http.HandlerFunc) {
	s.Router.Delete(path, handler)
}

func (s *WebServer) Start() error {
	return http.ListenAndServe(s.WebServerPort, s.Router)
}
