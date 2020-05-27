package server

import (
	"net/http"

	"cloud.bsdrocker.com/CrowderSoup/di-example/config"

	"go.uber.org/zap"
)

// Server our HTTP server
type Server struct {
	Address string
	Logger  *zap.SugaredLogger
}

// NewServer returns a new server
func NewServer(config *config.Config, logger *zap.SugaredLogger) *Server {
	server := &Server{
		Address: config.Address,
		Logger:  logger,
	}

	server.initRoutes()

	return server
}

func (s *Server) initRoutes() {
	http.HandleFunc("/", s.hello)
}

func (s *Server) hello(w http.ResponseWriter, r *http.Request) {
	s.Logger.Infow("recieved request",
		"RequestURI", r.RequestURI,
	)
	w.WriteHeader(200)
	w.Write([]byte("Hello World"))
}

// Run starts our HTTP server
func (s *Server) Run() error {
	s.Logger.Infow("starting http server",
		"Address", s.Address,
	)

	return http.ListenAndServe(s.Address, http.DefaultServeMux)
}
