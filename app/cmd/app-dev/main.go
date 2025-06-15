package main

import (
	"log"
	"miniservice/app/internal/adapter/handler"
	"miniservice/app/internal/adapter/repository"
	"miniservice/app/internal/adapter/router"
	"miniservice/app/internal/config"
	"miniservice/app/internal/domain/service"
	"miniservice/app/internal/persistence/database"
	"miniservice/app/pkg/jwtx"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.LoadENV()

	database.InitializeConnect()

	repository := repository.New(database.Initialize)
	service := service.New(repository)
	handler := handler.New(service)

	jwtx := jwtx.New()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-signalChan
		log.Printf("Received signal: %s. Shutting down...", sig)
		database.Close()
		os.Exit(0)
	}()

	conf := &router.ConfigVariantRouter{
		Handler: handler,
		Jwtx:    jwtx,
	}

	router := router.New(conf)
	router.Register()

	database.Close()
}
