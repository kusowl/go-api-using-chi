package main

import (
	"context"
	"log/slog"
	"os"
)

func main() {
	ctx := context.Background()

	dbConfig := dbConfig{
		connection: "pgsql",
		host:       "localhost",
		port:       5432,
		database:   "postgres",
		username:   "postgres",
		password:   "password",
	}

	config := config{
		address:  "localhost",
		port:     8080,
		dbConfig: dbConfig,
	}

	app := Application{
		config: config,
	}

	if err := app.run(ctx, app.mount()); err != nil {
		slog.Error("Application is shutting down to error: ", "error", err)
		os.Exit(1)
	}
}
