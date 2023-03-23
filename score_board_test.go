package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_(t *testing.T) {
	cases := []struct {
		name     string
		expected ScoreBoard
		start    ScoreBoard
		call     func(ScoreBoard) ScoreBoard
	}{
		{
			name: "zero values equal each other",
			call: func(s ScoreBoard) ScoreBoard {
				return ScoreBoard{}
			},
		},
		{
			name:     "adding 1 to score",
			start:    ScoreBoard{Score: 0, Level: 1, Lines: 0},
			expected: ScoreBoard{Score: 1, Level: 1, Lines: 1},
			call: func(s ScoreBoard) ScoreBoard {
				return s.Add(1)
			},
		},
		{
			name:     "score changes from 9 to 10",
			start:    ScoreBoard{Score: 9, Level: 1, Lines: 9},
			expected: ScoreBoard{Score: 10, Level: 2, Lines: 10},
			call: func(s ScoreBoard) ScoreBoard {
				return s.Add(1)
			},
		},
		{
			name:     "score level changes from 2 to 3",
			start:    ScoreBoard{Score: 27, Level: 3, Lines: 27},
			expected: ScoreBoard{Score: 33, Level: 4, Lines: 29},
			call: func(s ScoreBoard) ScoreBoard {
				return s.Add(2)
			},
		},
		{
			name:     "scored 4 lines by level",
			start:    ScoreBoard{Score: 28, Level: 3, Lines: 21},
			expected: ScoreBoard{Score: 40, Level: 5, Lines: 25},
			call: func(s ScoreBoard) ScoreBoard {
				return s.Add(4)
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expected, c.call(c.start))
		})
	}
}
