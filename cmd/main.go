package main

import (
	"context"
	"log/slog"
	"os"

	"shop/internal/configuration"
)

func main() {
	ctx := context.Background()
	configuration.LoadAncestorDotEnv()

	cfg := config{
		addr: ":8080",
	}

	logger := slog.New((slog.NewTextHandler(os.Stdout, nil)))

	slog.SetDefault(logger)

	conn, err := configuration.ConnectPostgres(ctx)

	if err != nil {
		slog.Error("DB has failed to start with", "error", err)
		os.Exit(1)
	}

	defer conn.Close()

	logger.Info("postgres: connected")

	api := &application{
		config: cfg,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("Server has failed to start with", "error", err)
		os.Exit(1)
	}
}
