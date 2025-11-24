package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/AntonTsoy/review-pull-request-service/internal/config"
	"github.com/AntonTsoy/review-pull-request-service/internal/database"
	"github.com/AntonTsoy/review-pull-request-service/internal/transport/http/handlers"
	"github.com/AntonTsoy/review-pull-request-service/internal/transport/http/server"
	"github.com/AntonTsoy/review-pull-request-service/internal/repository"
	"github.com/AntonTsoy/review-pull-request-service/internal/service"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	db, err := database.New(ctx, cfg)
	if err != nil {
		log.Fatalf("failed to create DB instance: %v", err)
	}
	defer db.Close()

	if err = db.HealthCheck(ctx); err != nil {
		log.Fatalf("failed to open connection with database: %v", err)
	}

	log.Println("DB connection opened!")

	repository := repository.NewRepository()

	service := service.NewService(db, repository)

	handlers := handlers.NewHandlers(service)

	server := server.New(handlers)
	_ = server

	// TODO: запуск сервера в отдельной горутине

	<-ctx.Done()
	log.Println("Shutting down...")
}
