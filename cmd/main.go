package main

import (
	"lecture/internal/data"
	"lecture/internal/http/handlers"
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
	http.HandleFunc("/", handler.PhraseOfTheDay)

	panic(http.ListenAndServe(":3333", nil))
}
