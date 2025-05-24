package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"github.com/go-chi/chi"
)

type Server struct {
	Log        *log.Logger
	HttpServer *http.Server
}

func NewServer(log *log.Logger) *Server {
	r := chi.NewRouter()
	r.Get("/", handlers.IndexHandler)
	r.Post("/upload", handlers.MainHandler)
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Start server error: %s", err.Error())

	}

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ErrorLog:     log,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	return &Server{Log: log, HttpServer: httpServer}
}
