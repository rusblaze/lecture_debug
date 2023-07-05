package main

import (
	"github.com/rs/zerolog/log"
	"lecture/internal/data"
	"lecture/internal/http/handlers"
	"lecture/pkg/http/interceptor"
	"net/http"
	"os"
)

func main() {
	file, err := os.Open("data/phrases.data")
	if err != nil {
		panic(err)
	}
	repo := data.NewFilePhraseRepository(file)
	handler := handlers.PhraseOfTheDayHandler{PhrasesRepo: repo}
	http.HandleFunc("/", interceptor.RecoverHandler(interceptor.LogHandler(handler.PhraseOfTheDay)))
	log.Info().Msg("Starting HTTP server")
	if err := http.ListenAndServe(":3333", nil); err != nil {
		log.Fatal().Err(err).Msg("")
	}
	log.Info().Msg("Exit HTTP server")
}
