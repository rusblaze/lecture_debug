package data

import (
	"bufio"
	"context"
	"encoding/csv"
	"lecture/domain"
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

func (repo *filePhraseRepository) GetPhraseOfTheDay(ctx context.Context) *domain.Phrase {
	fileScanner := bufio.NewScanner(repo.storage)
	fileScanner.Split(bufio.ScanLines)

	randomLineNumber := rand.Int31n(50)
	var lineNumber int32
	fileScanner.Scan()
	line := fileScanner.Text()
	for lineNumber = 1; lineNumber < randomLineNumber && fileScanner.Scan(); lineNumber++ {
		line = fileScanner.Text()
	}
	r := csv.NewReader(strings.NewReader(line))
	r.Comma = '\t'
	record, _ := r.Read()
	return domain.NewPhrase(record[0], record[1])
}
