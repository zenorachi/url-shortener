package tests

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/zenorachi/url-shortener/pkg/api"
	"github.com/zenorachi/url-shortener/tests/mocks"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"testing"
)

func TestShortenMethod(t *testing.T) {
	mockServer := &mocks.MockServer{}
	mockResponse := &api.ShortenResponse{ShortedUrl: "mockedShortenedUrl"}
	mockServer.On("Shorten", mock.Anything, mock.Anything).Return(mockResponse, nil)

	grpcServer := grpc.NewServer()

	api.RegisterUrlShortenerServer(grpcServer, mockServer)

	lis, err := net.Listen("tcp", "localhost:0")
	require.NoError(t, err, "Failed to listen")

	go func() {
		_ = grpcServer.Serve(lis)
	}()
	defer grpcServer.Stop()

	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err, "Failed to dial server")
	defer conn.Close()

	client := api.NewUrlShortenerClient(conn)

	shortenRequest := &api.ShortenRequest{
		OriginalUrl: "https://example.com",
	}
	shortenResponse, err := client.Shorten(context.Background(), shortenRequest)
	require.NoError(t, err, "Failed to call Shorten")

	assert.Equal(t, mockResponse.ShortedUrl, shortenResponse.ShortedUrl)
}
