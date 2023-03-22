package main

import "github.com/lcaballero/ebiten-01/shapes"

// ScoreBoard records the current values that are shown on the board
// during the game
type ScoreBoard struct {
	Score int
	Lines int
	Level int
}

func (s ScoreBoard) Add(n int) ScoreBoard {
	score := s.Score + (s.Level * n)
	level := (score / 10) + 1
	return ScoreBoard{
		Score: score,
		Lines: s.Lines + n,
		Level: level,
	}
}

// Velocity reports the vertical speed of falling blocks for the
// current level
func (s ScoreBoard) Velocity() shapes.Vec {
	return shapes.Vec{0, 5 * float64(s.Level+1)}
}
