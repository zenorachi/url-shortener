package main

import (
	"fmt"
	"github.com/zenorachi/url-shortener/internal/config"
	_ "github.com/zenorachi/url-shortener/pkg/api"
	"log/slog"
)

const (
	dir  = "./configs"
	file = "main"
)

func main() {
	config.InitENV(".env")
	cfg, err := config.New(dir, file)
	if err != nil {
		slog.Error("config", err)
	}
	fmt.Println(cfg)
}
