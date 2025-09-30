package main

import (
	"log/slog"
	"os"

	"github.com/kyomel/pos-app/internal/config"
	"github.com/kyomel/pos-app/internal/server"
)

func main() {
	err := config.LoadConfig("config.yml")
	if err != nil {
		panic(err)
	}

	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	log := slog.New(logHandler)
	slog.SetDefault(log)

	if err := server.Start(); err != nil {
		panic(err)
	}
}
