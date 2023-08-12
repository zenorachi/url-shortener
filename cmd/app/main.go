package main

import (
	"github.com/zenorachi/url-shortener/internal/config"
	"github.com/zenorachi/url-shortener/internal/repository"
	"github.com/zenorachi/url-shortener/internal/service"
	"github.com/zenorachi/url-shortener/internal/transport/grpc"
	"github.com/zenorachi/url-shortener/internal/transport/grpc/handler"
	_ "github.com/zenorachi/url-shortener/pkg/api"
	"github.com/zenorachi/url-shortener/pkg/base62"
	"github.com/zenorachi/url-shortener/pkg/db/redis"
	"log/slog"
)

const (
	configDir  = "./configs"
	configFile = "main"
	envFile    = ".env"
)

func main() {
	if err := config.InitENV(envFile); err != nil {
		slog.Error(".env file not found", err)
	}
	cfg, err := config.New(configDir, configFile)
	if err != nil {
		slog.Error("config", err)
	}

	client, err := redis.New(cfg.Redis)
	if err != nil {
		slog.Error("redis client", err)
	}

	repo := repository.NewUrls(client)

	b62Generator := base62.NewGenerator()

	urlsService := service.NewUrls(b62Generator, repo)

	grpcHandler := handler.NewHandler(urlsService)

	server := grpc.NewServer(grpcHandler)

	slog.Info("server started...")
	if err = server.ListenAndServe(cfg); err != nil {
		slog.Error("server error", err)
	}
}
