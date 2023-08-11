package handler

import "github.com/zenorachi/url-shortener/pkg/api"

type Url interface {
	Shorten(url string)
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
