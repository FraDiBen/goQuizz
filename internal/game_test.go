package internal

import (
	"github.com/stretchr/testify/assert"
	"quizz/internal/reader"
	"strings"
	"testing"
)

var example_csv = strings.NewReader(`5+5,10
7+3,10
1+1,2
8+3,11
1+2,3
8+6,14
3+1,4
1+4,5
5+1,6
2+3,5
3+3,6
2+4,6
5+2,7`)

type fakeResponder struct {
	offset  *int
	answers []string
}

func (f fakeResponder) ReadString(delim byte) (string, error) {
	if *f.offset >= len(f.answers) {
		zero := 0
		f.offset = &zero
	}
	a := f.answers[*f.offset]
	*f.offset++
	return a, nil
}

func TestGame_Play(t *testing.T) {

	t.Run("first game", func(t *testing.T) {
		zero := 0
		resp := fakeResponder{
			offset:  &zero,
			answers: []string{"10", "2", "11", "3"},
		}

		g := Game{
			reader:         reader.Read,
			inputProcessor: resp.ReadString,
		}

		assert.False(t, g.Play(example_csv))
	})

	t.Run("second game", func(t *testing.T) {
		var win_csv = strings.NewReader("5+5,10\n7+3,10\n1+1,2")
		zero := 0
		resp := fakeResponder{
			offset:  &zero,
			answers: []string{"10", "10", "2"},
		}
		g := Game{
			reader:         reader.Read,
			inputProcessor: resp.ReadString,
		}

		assert.True(t, g.Play(win_csv))
	})

}
