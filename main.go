package main

import (
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func setupRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthCheckHandler)
	return mux
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	router := setupRouter()

	log.Info().Msg("Service is healthy and ready to deploy")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal().Err(err).Msg("Server failed to start")
	}
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("healthy"))
}
