package server

import (
	"context"
	"net/http"
	"time"

	"github.com/CrowderSoup/di-example/config"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Server our HTTP server
type Server struct {
	Address string
	Logger  *zap.SugaredLogger
	mux     *http.ServeMux
	http    *http.Server
}

// NewServer returns a new server
func NewServer(config *config.Config, logger *zap.SugaredLogger, mux *http.ServeMux) *Server {
	server := &Server{
		Address: config.Address,
		Logger:  logger,
		mux:     mux,
	}

	server.initRoutes()

	s := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.http = s

	return server
}

func (s *Server) initRoutes() {
	s.mux.HandleFunc("/", s.hello)
}

func (s *Server) hello(w http.ResponseWriter, r *http.Request) {
	s.Logger.Infow("recieved request",
		"RequestURI", r.RequestURI,
	)
	w.WriteHeader(200)
	w.Write([]byte("Hello World"))
}

// Run starts our HTTP server
func Run(lc fx.Lifecycle, s *Server) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				s.Logger.Infow("starting http server",
					"Address", s.Address,
				)

				go s.http.ListenAndServe()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				s.http.Shutdown(ctx)
				return nil
			},
		},
	)
}

// Module provided to fx
var Module = fx.Options(
	fx.Provide(
		http.NewServeMux,
		NewServer,
	),
)
