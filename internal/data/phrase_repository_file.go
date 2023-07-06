package data

import (
	"bufio"
	"context"
	"encoding/csv"
	"lecture/domain"
	"lecture/pkg/log"
	"math/rand"
	"os"
	"strings"
)

type filePhraseRepository struct {
	storage *os.File
}

func NewFilePhraseRepository(storage *os.File) domain.PhrasesRepository {
	return &filePhraseRepository{storage: storage}
}

func (repo *filePhraseRepository) GetPhraseOfTheDay(ctx context.Context) (*domain.Phrase, error) {
	repo.storage.Seek(0, 0) // <- Устанавливаем смещение в начало файла
	fileScanner := bufio.NewScanner(repo.storage)
	fileScanner.Split(bufio.ScanLines)

	randomLineNumber := rand.Int31n(50)
	var lineNumber int32
	fileScanner.Scan()
	line := fileScanner.Text()
	log.Debug(ctx).Msgf("randomLineNumber = %d", randomLineNumber)
	for lineNumber = 1; lineNumber < randomLineNumber && fileScanner.Scan(); lineNumber++ {
		line = fileScanner.Text()
	}
	log.Debug(ctx).Msgf("line = %s, lineNumber = %d", line, lineNumber)
	r := csv.NewReader(strings.NewReader(line))
	r.Comma = '\t'
	record, err := r.Read()
	if err != nil {
		log.Error(ctx).Err(err).
			Str("line", line).
			Int32("randomLineNumber", randomLineNumber).
			Int32("lineNumber", lineNumber).
			Msg("error reading csv line")
		return nil, err
	}
	return domain.NewPhrase(record[0], record[1]), nil
}
