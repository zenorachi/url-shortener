package handler

import (
	"context"
	"github.com/zenorachi/url-shortener/pkg/api"
)

type Url interface {
	Shorten(ctx context.Context, request *api.ShortenRequest) (*api.ShortenResponse, error)
	Redirect(ctx context.Context, request *api.RedirectRequest) (*api.RedirectResponse, error)
}

type Handler struct {
	api.UnimplementedUrlShortenerServer
	url Url
}

func NewHandler(url Url) *Handler {
	return &Handler{
		url: url,
	}
}

func (h *Handler) Shorten(ctx context.Context, request *api.ShortenRequest) (*api.ShortenResponse, error) {
	return h.url.Shorten(ctx, request)
}

func (h *Handler) Redirect(ctx context.Context, request *api.RedirectRequest) (*api.RedirectResponse, error) {
	return h.url.Redirect(ctx, request)
}
