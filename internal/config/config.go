package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
	"github.com/zenorachi/url-shortener/pkg/db/redis"
)

type Config struct {
	GRPCServer struct {
		Network string
		Port    string
	}

	GatewayServer struct {
		Network string
		Host    string
		Port    string
	}

	Redis redis.ClientConfig
}

func InitENV(filename string) error {
	return godotenv.Load(filename)
}

func New(directory, filename string) (*Config, error) {
	cfg := new(Config)

	viper.AddConfigPath(directory)
	viper.SetConfigName(filename)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	if err := envconfig.Process("redis", &cfg.Redis); err != nil {
		return nil, err
	}

	return cfg, nil
}
