package repository

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/zenorachi/url-shortener/model"
)

type Urls struct {
	client *redis.Client
}

func NewUrls(client *redis.Client) *Urls {
	return &Urls{client: client}
}

func (u *Urls) Create(ctx context.Context, url model.URL) error {
	_, err := u.client.Set(ctx, url.Shortened, url.Original, 0).Result() //todo ttl
	if err != nil {
		return err
	}

	return nil
}

func (u *Urls) GetByShorted(ctx context.Context, shorted string) (string, error) {
	original, err := u.client.Get(ctx, shorted).Result()
	if err != nil {
		return "", err
	}

	return original, nil
}
