package app

import (
	"github.com/zenorachi/url-shortener/internal/config"
	"github.com/zenorachi/url-shortener/internal/repository"
	"github.com/zenorachi/url-shortener/internal/service"
	"github.com/zenorachi/url-shortener/internal/transport/grpc"
	"github.com/zenorachi/url-shortener/internal/transport/grpc/handler"
	_ "github.com/zenorachi/url-shortener/pkg/api"
	"github.com/zenorachi/url-shortener/pkg/base62"
	"github.com/zenorachi/url-shortener/pkg/db/redis"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	configDir  = "./configs"
	configFile = "main"
	envFile    = ".env"
)

func Run() error {
	if err := config.InitENV(envFile); err != nil {
		slog.Error(".env file")
		return err
	}

	cfg, err := config.New(configDir, configFile)
	if err != nil {
		slog.Error("config")
		return err
	}

	client, err := redis.New(cfg.Redis)
	if err != nil {
		slog.Error("redis client")
		return err
	}

	repo := repository.NewUrls(client)

	b62Generator := base62.NewGenerator()

	urlsService := service.NewUrls(b62Generator, repo)

	grpcHandler := handler.NewHandler(urlsService)

	server := grpc.NewServer(grpcHandler)

	slog.Info("server started...")
	go func() {
		if err = server.ListenAndServe(cfg.Server.Network, cfg.Server.Port); err != nil {
			log.Fatalln("error starting server", err)
		}
	}()
	slog.Info("server started...")

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM)

	<-quit
	slog.Info("server shutting down")

	server.Stop()
	return nil
}
