package server

import (
	"context"
	"errors"
	"github.com/rs/cors"
	"log"
	"net/http"
	"rest-websockets-go/database"
	"rest-websockets-go/repository"
	"rest-websockets-go/websockets"

	"github.com/gorilla/mux"
)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseUrl string
}

type Server interface {
	Config() *Config
	Hub() *websockets.Hub
}

type Broker struct {
	config *Config
	router *mux.Router
	hub    *websockets.Hub
}

func (b *Broker) Config() *Config {
	return b.config
}

func (b *Broker) Hub() *websockets.Hub {
	return b.hub
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}
	if config.JWTSecret == "" {
		return nil, errors.New("jwt secret is required")
	}
	if config.DatabaseUrl == "" {
		return nil, errors.New("database url is required")
	}
	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
		hub:    websockets.NewHub(),
	}
	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)
	handler := cors.Default().Handler(b.router)
	// Connecting to Database
	repo, err := database.NewPostgresRepository(b.config.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	go b.hub.Run()
	repository.SetRepository(repo)
	log.Println("starting server on port", b.config.Port)
	if err := http.ListenAndServe(b.config.Port, handler); err != nil {
		log.Println("error starting server:", err)
	} else {
		log.Fatalf("server stopped")
	}
}
