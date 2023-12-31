package grpc

import (
	"net"

	"github.com/zenorachi/url-shortener/pkg/api"
	"google.golang.org/grpc"
)

type Server struct {
	server  *grpc.Server
	handler api.UrlShortenerServer
}

func NewServer(handler api.UrlShortenerServer) *Server {
	return &Server{
		server:  grpc.NewServer(),
		handler: handler,
	}
}

func (s *Server) ListenAndServe(network, port string) error {
	lis, err := net.Listen(network, ":"+port)
	if err != nil {
		return err
	}

	api.RegisterUrlShortenerServer(s.server, s.handler)

	if err = s.server.Serve(lis); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() {
	s.server.GracefulStop()
}
