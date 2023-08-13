package gateway

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"github.com/zenorachi/url-shortener/internal/transport/grpc/handler"
	"github.com/zenorachi/url-shortener/pkg/api"
	"net"
	"net/http"
)

type Server struct {
	grpcHandler *handler.Handler
	grpcMux     *runtime.ServeMux
	ctx         context.Context
}

func NewServer(grpcHandler *handler.Handler, ctx context.Context) *Server {
	return &Server{
		grpcHandler: grpcHandler,
		grpcMux:     runtime.NewServeMux(),
		ctx:         ctx,
	}
}

func (s *Server) ListenAndServe(network, host, port string) error {
	addr := fmt.Sprintf("%s:%s", host, port)

	lis, err := net.Listen(network, addr)
	if err != nil {
		return err
	}

	err = api.RegisterUrlShortenerHandlerServer(s.ctx, s.grpcMux, s.grpcHandler)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/", s.grpcMux)

	prefix := "/docs/"
	fileServer := httpSwagger.Handler(
		httpSwagger.URL("/docs/url-shortener.swagger.json"),
	)
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))

	if err = http.Serve(lis, mux); err != nil {
		return err
	}

	return nil
}
