package main

import (
	"net/http"

	"github.com/koding/multiconfig"
	"go.uber.org/zap"
)

// Config our app config
type Config struct {
	Address string `default:":3000"`
}

func main() {
	var config Config
	m := multiconfig.New()
	m.MustLoad(&config)

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sugar.Infow("recieved request",
			"RequestURI", r.RequestURI,
		)
		w.WriteHeader(200)
		w.Write([]byte("Hello World"))
	})

	sugar.Infow("starting http server",
		"Address", config.Address,
	)
	sugar.Fatal(http.ListenAndServe(config.Address, http.DefaultServeMux))
}
