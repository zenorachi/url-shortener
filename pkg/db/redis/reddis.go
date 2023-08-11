package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type ClientConfig struct {
	Host     string
	Port     string
	Password string
	DBIndex  int
}

func New(cfg ClientConfig) (*redis.Client, error) {
	ctx := context.Background()

	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       cfg.DBIndex,
	})

	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return client, nil
}
