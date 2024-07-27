package server

import (
	"context"
	"net/http"
	"time"

	"github.com/SurkovIlya/message-handler-app/internal/model"
)

type Broker interface {
	Receive(message model.Message) error
}

type StatCollector interface {
	GetStat() (model.Statistic, error)
}

type Server struct {
	httpServer *http.Server
	Messager   Broker
	SC         StatCollector
}

func New(port string, msgr Broker, sc StatCollector) *Server {
	s := &Server{
		httpServer: &http.Server{
			Addr:           ":" + port,
			MaxHeaderBytes: 1 << 20, // 1 MB
			ReadTimeout:    5000 * time.Millisecond,
			WriteTimeout:   5000 * time.Millisecond,
		},
		Messager: msgr,
		SC:       sc,
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
	mux.HandleFunc("GET /v1/getstatistics", s.GetStatistics)

	s.httpServer.Handler = mux
}
