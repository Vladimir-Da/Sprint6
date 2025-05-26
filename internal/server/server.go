package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Log        *log.Logger
	HttpServer *http.Server
}

func NewServer(Logger *log.Logger) *Server {
	mux := http.NewServeMux()
	// хендл функции сервера
	mux.HandleFunc("GET /", handlers.IndexHandler)
	mux.HandleFunc("POST /upload", handlers.MainHandler)

	return &Server{
		Log: Logger,
		HttpServer: &http.Server{
			Addr:         ":8080",
			Handler:      mux,
			ErrorLog:     Logger,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	return s.HttpServer.ListenAndServe()
}
