package reader

import (
	"github.com/stretchr/testify/assert"
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

func TestCSV_Read(t *testing.T) {

	t.Run("test read", func(t *testing.T) {
		questions := example_csv
		q := Read(questions)
		assert.Greater(t, len(q), 1)
	})

}
func TestCSV_Parsing(t *testing.T) {

	t.Run("test parse a single record", func(t *testing.T) {
		record_fields := []string{"2+4", "6"}

		q, err := parse(record_fields)

		assert.NoError(t, err)
		assert.Equal(t, 2, q.A)
		assert.Equal(t, 4, q.B)
		assert.Equal(t, 6, q.Answer.Correct)
	})

}
