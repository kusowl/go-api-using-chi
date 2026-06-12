package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Application struct {
	config config
}

func (app Application) mount() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.ClientIPFromRemoteAddr)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(time.Minute))

	router.Get("/health", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Good"))
	})

	return router
}

func (app Application) run(handler http.Handler) error {
	slog.Info("Application running....")

	return http.ListenAndServe(fmt.Sprintf("%s:%d", app.config.address, app.config.port), handler)
}

type config struct {
	address string
	port    int32
}
