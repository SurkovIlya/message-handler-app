package server

import (
	"context"
	"net/http"
	"time"
)

type MessagerStorage interface {
	Receiving(message ReceivingBody) error
}

type Server struct {
	httpServer *http.Server
	Messager   MessagerStorage
}

func New(port string, msgr MessagerStorage) *Server {
	s := &Server{
		httpServer: &http.Server{
			Addr:           ":" + port,
			MaxHeaderBytes: 1 << 20, // 1 MB
			ReadTimeout:    200 * time.Millisecond,
			WriteTimeout:   200 * time.Millisecond,
		},
		Messager: msgr,
	}

	s.initRoutes()

	return s
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) initRoutes() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /v1/receiving", s.ReceivingMessages)

	s.httpServer.Handler = mux
}
