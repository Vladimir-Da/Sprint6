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

func NewServer(log *log.Logger) *Server {
	//r := chi.NewRouter()
	//r.Get("/", handlers.IndexHandler)
	//r.Post("/upload", handlers.MainHandler)
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", handlers.MainHandler)
	mux.HandleFunc("/", handlers.IndexHandler)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     log,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	return &Server{Log: log, HttpServer: httpServer}
}
