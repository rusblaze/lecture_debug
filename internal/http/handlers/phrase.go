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
	phrase, err := h.PhrasesRepo.GetPhraseOfTheDay(r.Context())
	if err != nil {
		http.Error(w, "Произошла ошибка при обработке запроса. Попробуйте еще раз", http.StatusInternalServerError)
		return
	}
	buff := bytes.NewBufferString(fmt.Sprintf("%s\n%s", phrase.GetPhrase(), phrase.GetAuthor()))
	w.Write(buff.Bytes())
}
