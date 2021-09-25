package reader

import (
	"errors"
	"io"
	"log"
	"quizz/internal/quiz"
	"strconv"
	"strings"
)
import "encoding/csv"

func Read(example_csv io.Reader) []*quiz.Question {
	questions := make([]*quiz.Question, 0)
	csvReader := csv.NewReader(example_csv)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("error reading file", err)
		}
		question, err := parse(rec)
		if err != nil {
			continue
		}
		questions = append(questions, question)
	}
	return questions
}

func parse(record_fields []string) (*quiz.Question, error) {
	if len(record_fields) < 2 {
		return nil, errors.New("not enough fields")
	}
	ab := strings.Split(record_fields[0], "+")
	a, _ := strconv.ParseInt(ab[0], 10, 32)
	b, _ := strconv.ParseInt(ab[1], 10, 32)
	corr, _ := strconv.ParseInt(record_fields[1], 10, 32)
	return &quiz.Question{
		A: int(a),
		B: int(b),
		Answer: &quiz.Answer{
			Correct: int(corr),
			Given:   nil,
		},
	}, nil
}
