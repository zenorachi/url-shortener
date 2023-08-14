package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/zenorachi/url-shortener/pkg/api"
)

type MockServer struct {
	api.UnimplementedUrlShortenerServer
	mock.Mock
}

func (m *MockServer) Shorten(ctx context.Context, req *api.ShortenRequest) (*api.ShortenResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*api.ShortenResponse), args.Error(1)
}

func (m *MockServer) Redirect(ctx context.Context, req *api.RedirectRequest) (*api.RedirectResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*api.RedirectResponse), args.Error(1)
}
