package grpc

import (
	"github.com/zenorachi/url-shortener/pkg/api"
	"google.golang.org/grpc"
	"net"
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

func (s *Server) ListenAndServe(port string) error {
	lis, err := net.Listen("tcp", ":"+port)
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
