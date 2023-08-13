package app

import (
	"context"
	"github.com/zenorachi/url-shortener/internal/config"
	"github.com/zenorachi/url-shortener/internal/repository"
	"github.com/zenorachi/url-shortener/internal/service"
	"github.com/zenorachi/url-shortener/internal/transport/gateway"
	"github.com/zenorachi/url-shortener/internal/transport/grpc"
	"github.com/zenorachi/url-shortener/internal/transport/grpc/handler"
	_ "github.com/zenorachi/url-shortener/pkg/api"
	"github.com/zenorachi/url-shortener/pkg/base62"
	"github.com/zenorachi/url-shortener/pkg/db/redis"
	"github.com/zenorachi/url-shortener/pkg/logger"
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
		logger.ErrorService("config", err)
		return err
	}

	cfg, err := config.New(configDir, configFile)
	if err != nil {
		logger.ErrorService("config", err)
		return err
	}

	client, err := redis.New(cfg.Redis)
	if err != nil {
		logger.ErrorService("redis", err)
		return err
	}

	repo := repository.NewUrls(client)

	b62Generator := base62.NewGenerator()

	urlsService := service.NewUrls(b62Generator, repo)

	grpcHandler := handler.NewHandler(urlsService)

	grpcServer := grpc.NewServer(grpcHandler)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gatewayServer := gateway.NewServer(grpcHandler, ctx)

	slog.Info("gateway server started...")
	go func() {
		err = gatewayServer.
			ListenAndServer(cfg.GatewayServer.Network, cfg.GatewayServer.Host, cfg.GatewayServer.Port)
		if err != nil {
			log.Fatalln("error starting gateway server", err)
		}
	}()

	slog.Info("grpc server started...")

	go func() {
		if err = grpcServer.ListenAndServe(cfg.GRPCServer.Network, cfg.GRPCServer.Port); err != nil {
			log.Fatalln("error starting grpc server", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM)

	<-quit
	slog.Info("server shutting down")

	grpcServer.Stop()
	return nil
}
