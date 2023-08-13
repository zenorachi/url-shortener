package main

import (
	"github.com/zenorachi/url-shortener/internal/app"
	_ "github.com/zenorachi/url-shortener/pkg/api"
	"log/slog"
)

const (
	configDir  = "./configs"
	configFile = "main"
	envFile    = ".env"
)

func main() {
	if err := app.Run(); err != nil {
		slog.Error("error starting server", err)
	}
}
