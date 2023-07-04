package handlers

import (
	"bytes"
	"fmt"
	"lecture/domain"
	"net/http"
)

type PhraseOfTheDayHandler struct {
	PhrasesRepo domain.PhrasesRepository
}

func (h *PhraseOfTheDayHandler) PhraseOfTheDay(w http.ResponseWriter, r *http.Request) {
	phrase := h.PhrasesRepo.GetPhraseOfTheDay(r.Context())
	buff := bytes.NewBufferString(fmt.Sprintf("%s\n%s", phrase.GetPhrase(), phrase.GetAuthor()))
	w.Write(buff.Bytes())
}
