package service

import (
	"context"
	"github.com/zenorachi/url-shortener/model"
	"github.com/zenorachi/url-shortener/pkg/api"
)

type ShortCodeGenerator interface {
	GenerateCode() string
}

type UrlRepository interface {
	Create(ctx context.Context, url model.URL) error
	GetByShorted(ctx context.Context, shorted string) (string, error)
}

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
	shortUrl := u.generator.GenerateCode()
	url := model.NewUrl(request.OriginalUrl, shortUrl)

	if err := u.urlsRepo.Create(ctx, url); err != nil {
		return nil, err
	}
	return &api.ShortenResponse{ShortedUrl: shortUrl}, nil
}

func (u *Urls) GetByShorted(ctx context.Context, request *api.ShortedRequest) (*api.ShortedResponse, error) {
	url, err := u.urlsRepo.GetByShorted(ctx, request.ShortedUrl)
	if err != nil {
		return nil, err
	}
	return &api.ShortedResponse{OriginalUrl: url}, nil
}
