package main

import (
	"log/slog"

	"github.com/zenorachi/url-shortener/internal/app"
	_ "github.com/zenorachi/url-shortener/pkg/api"
)

func main() {
	if err := app.Run(); err != nil {
		slog.Error("unexpected error", "err", err)
	}
}
