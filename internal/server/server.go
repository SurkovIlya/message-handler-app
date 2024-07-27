package server

import (
	"context"
	"net/http"
	"time"

	"github.com/SurkovIlya/message-handler-app/internal/model"
)

type MessagerStorage interface {
	Receiving(message model.Message) error
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
			ReadTimeout:    1000 * time.Millisecond,
			WriteTimeout:   1000 * time.Millisecond,
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
	mux.HandleFunc("POST /vq/getstatistics", s.GetStatistics)

	s.httpServer.Handler = mux
}
