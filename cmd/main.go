package main

import (
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"lecture/internal/data"
	"lecture/internal/http/handlers"
	"lecture/internal/http/handlers/infra"
	"lecture/internal/tracing"
	"lecture/pkg/http/interceptor"
	"net/http"
	_ "net/http/pprof" // Magic is here
	"os"
)

// source cmd/main.go

func main() {
	file, err := os.Open("data/phrases.data")
	if err != nil {
		panic(err)
	}

	log.Logger = log.With().Caller().Timestamp().Logger()

	if err := tracing.SetupGlobalTracer("phrases", "0.0.2"); err != nil {
		log.Error().Err(err).Msg("")
	}

	repo := data.NewFilePhraseRepository(file)
	handler := handlers.PhraseOfTheDayHandler{PhrasesRepo: repo}
	phraseHandler := otelhttp.NewHandler(
		http.HandlerFunc(
			interceptor.RecoverHandler(
				interceptor.LogHandler(handler.PhraseOfTheDay))), "phrase-of-the-day")
	http.Handle("/", phraseHandler)
	http.HandleFunc("/infra/log_level", interceptor.RecoverHandler(interceptor.LogHandler(infra.SetLogLevel)))

	log.Info().Msg("Starting HTTP server")
	if err := http.ListenAndServe(":3333", nil); err != nil {
		log.Fatal().Err(err).Msg("")
	}
	log.Info().Msg("Exit HTTP server")
}
