package server

import (
	"context"
	"net/http"

	"cloud.bsdrocker.com/CrowderSoup/di-example/internal/config"
	"cloud.bsdrocker.com/CrowderSoup/di-example/internal/logging"

	"go.uber.org/fx"
)

// Handler for http requests
type Handler struct {
	mux *http.ServeMux
}

// NewHandler http handler
func NewHandler(s *http.ServeMux) *Handler {
	h := Handler{
		mux: s,
	}
	h.registerRoutes()

	return &h
}

// RegisterRoutes for all http endpoints
func (h *Handler) registerRoutes() {
	h.mux.HandleFunc("/", h.hello)
}

func (h *Handler) hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello World"))
}

// InvokeServer starts our http server
func InvokeServer(lc fx.Lifecycle, cfg *config.Config, handler *Handler, logger *logging.Logger) {
	lc.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go http.ListenAndServe(cfg.HTTPAddress, handler.mux)
				return nil
			},
			OnStop: func(context.Context) error {
				return nil
			},
		},
	)
}

// Module provided to fx
var Module = fx.Options(
	fx.Provide(
		http.NewServeMux,
		NewHandler,
	),
)
