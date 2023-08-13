package service

import (
	"context"
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/redis/go-redis/v9"
	"github.com/zenorachi/url-shortener/model"
	"github.com/zenorachi/url-shortener/pkg/api"
)

const urlLen = 10

type (
	ShortCodeGenerator interface {
		GenerateCode(len int) string
	}

	UrlRepository interface {
		Create(ctx context.Context, url model.URL) error
		GetByShorted(ctx context.Context, shorted string) (string, error)
	}
)

type Urls struct {
	generator ShortCodeGenerator
	urlsRepo  UrlRepository
}

func NewUrls(generator ShortCodeGenerator, repo UrlRepository) *Urls {
	return &Urls{
		generator: generator,
		urlsRepo:  repo,
	}
}

func (u *Urls) Shorten(ctx context.Context, request *api.ShortenRequest) (*api.ShortenResponse, error) {
	if !govalidator.IsURL(request.OriginalUrl) {
		return nil, errors.New("url is not valid")
	}

	shortUrl := u.generator.GenerateCode(urlLen)
	url := model.NewUrl(request.OriginalUrl, shortUrl)

	if err := u.urlsRepo.Create(ctx, url); err != nil {
		return nil, err
	}
	return &api.ShortenResponse{ShortedUrl: shortUrl}, nil
}

func (u *Urls) Redirect(ctx context.Context, request *api.RedirectRequest) (*api.RedirectResponse, error) {
	url, err := u.urlsRepo.GetByShorted(ctx, request.ShortedUrl)
	if err != nil {
		if err != redis.Nil {
			return nil, err
		}
		return nil, errors.New("url not found")
	}

	return &api.RedirectResponse{OriginalUrl: url}, nil
}
