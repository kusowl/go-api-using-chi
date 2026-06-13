package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
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

func (app Application) run(ctx context.Context, handler http.Handler) error {
	app.connectDb(ctx)
	slog.Info("Application running....")

	return http.ListenAndServe(fmt.Sprintf("%s:%d", app.config.address, app.config.port), handler)
}

func (app Application) connectDb(ctx context.Context) {
	conn, err := pgx.Connect(ctx, app.config.dbConfig.dsn())
	if err != nil {
		slog.Error("Database connection failed!", "error", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)
}

type config struct {
	address  string
	port     int32
	dbConfig dbConfig
}

type dbConfig struct {
	connection string
	host       string
	port       int32
	database   string
	username   string
	password   string
}

func (db dbConfig) dsn() string {
	dsn := ""
	switch db.connection {
	case "pgsql":
		dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", db.host, db.port, db.username, db.password, db.database)
	default:
		slog.Warn("Unsupported database connection type", "connection", db.connection)
		return ""
	}
	return dsn
}
