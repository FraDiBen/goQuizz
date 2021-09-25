package internal

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"quizz/internal/quiz"
	"quizz/internal/reader"
	"strconv"
	"strings"
)

func NewCLIGame() *Game {
	r := bufio.NewReader(os.Stdin)
	return &Game{
		forecasted:     0,
		reader:         reader.Read,
		inputProcessor: r.ReadString,
	}
}

type Game struct {
	forecasted     int
	reader         func(io.Reader) []*quiz.Question
	inputProcessor func(byte) (string, error)
}

func (g Game) load(questionFile io.Reader) []*quiz.Question {
	return g.reader(questionFile)
}

func (g Game) Play(questionFile io.Reader) bool {
	questions := g.load(questionFile)
	for _, q := range questions {
		answer := g.ask(q)
		if !g.isCorrect(*answer.Given, q.Answer.Correct) {
			continue
		}
		g.forecasted++
	}
	return g.forecasted == len(questions)
}

func (g Game) ask(q *quiz.Question) quiz.Answer {
	fmt.Printf(">> Question %d + %d ?\n\tYour answer:  ", q.A, q.B)
	read, err := g.inputProcessor('\n')
	fmt.Println()
	if err != nil {
		return quiz.Answer{}
	}

	given, err := strconv.ParseInt(strings.Trim(read, "\n"), 10, 32)
	if err != nil {
		log.Println(err)
		return quiz.Answer{}
	}

	gint := int(given)
	return quiz.Answer{
		Correct: 0,
		Given:   &gint,
	}
}

func (g Game) isCorrect(given, correct int) bool {
	return given == correct
}
